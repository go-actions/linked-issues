package parser

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type HTMLParser struct {
}

func (parser *HTMLParser) Parse() ([]string, error) {
	return []string{}, nil
}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{}
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
