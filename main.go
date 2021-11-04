package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Input struct {
	prURL  string
	format string
	tag    string
	attr   html.Attribute
}

type OutputFormatter interface {
	Format([]string) string
}

const (
	IssueNumber      = "IssueNumber"
	IssueURL         = "IssueURL"
	ExternalIssueRef = "ExternalIssueRef"
)

var input Input

func init() {
	flag.StringVar(&input.prURL, "pr-url", "", "URL of the pull request.")
	flag.StringVar(&input.tag, "tag", "form", "HTML tag which contans the issue links.")
	flag.StringVar(&input.attr.Key, "attr-key", "aria-label", "The key of the tag attribute.")
	flag.StringVar(&input.attr.Val, "attr-val", "Link issues", "The value of the tag attribute.")
	flag.StringVar(&input.format, "format", IssueNumber, "Format of the output issue list.")
}

func main() {
	// Parse the flags
	flag.Parse()

	//  Find the linked issues
	issues, err := input.findLinkedIssues()
	if err != nil {
		fmt.Println("Failed to find the linked Issues. Reason: ", err.Error())
		os.Exit(1)
	}

	var formatter OutputFormatter

	switch input.format {
	case IssueNumber:
		formatter = &NumberFormatter{}
	case IssueURL:
		formatter = &URLFormatter{}
	case ExternalIssueRef:
		formatter = &ExternalIssueRefFormatter{}
	default:
		fmt.Println("Unknown format: ", input.format)

	}
	fmt.Printf("::set-output name=issues::%s\n", formatter.Format(issues))
}

func (i *Input) findLinkedIssues() ([]string, error) {
	if i.prURL == "" {
		return nil, fmt.Errorf("pr-url must not be empty")
	}

	// Download the HTML page of the PR
	response, err := http.Get(i.prURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}

	// Now, find the node that has the HTML tag with the desired attribute
	tag := i.findTagNode(doc)
	if tag == nil {
		return nil, fmt.Errorf("tag %q with attribute %s=%q does not exist", i.tag, i.attr.Key, i.attr.Val)
	}

	// Now, find the <a> tags which contains the issue URL.
	return findLinks(tag)
}

func (i *Input) findTagNode(n *html.Node) *html.Node {
	// If the desired node is found, return it.
	if n.Type == html.ElementNode && n.Data == i.tag {
		for _, a := range n.Attr {
			if a.Key == i.attr.Key && a.Val == i.attr.Val {
				return n
			}
		}
	}
	// Recursively look into the next nodes for the desired tag.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := i.findTagNode(c)
		if node != nil {
			return node
		}
	}
	return nil
}

func findLinks(tag *html.Node) ([]string, error) {
	links := make([]string, 0)

	// Declare an closure function which recursively look for links.
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	// Call the closure function.
	f(tag)

	if len(links) == 0 {
		return nil, fmt.Errorf("no links found inside the provided tag: %s", tag.Data)
	}
	return links, nil
}

type NumberFormatter struct{}

func (n *NumberFormatter) Format(issues []string) string {
	output := make([]string, 0)

	for i := range issues {
		parts := strings.Split(issues[i], "/")
		output = append(output, parts[len(parts)-1])
	}
	return strings.Join(output, " ")
}

type URLFormatter struct{}

func (n *URLFormatter) Format(issues []string) string {
	return strings.Join(issues, " ")
}

type ExternalIssueRefFormatter struct{}

func (n *ExternalIssueRefFormatter) Format(issues []string) string {
	output := make([]string, 0)

	for i := range issues {
		parts := strings.Split(issues[i], "/")
		output = append(output, fmt.Sprintf("%s/%s#%s", parts[len(parts)-4], parts[len(parts)-3], parts[len(parts)-1]))
	}
	return strings.Join(output, " ")
}
