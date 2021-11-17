package finder

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type htmlParser struct {
	prURL  string
	tag    string
	attr   html.Attribute
	client http.Client
}

func NewHTMLParser(prURL string, opts ...func(p *htmlParser)) *htmlParser {
	p := &htmlParser{
		prURL: prURL,
		tag:   "form",
		attr: html.Attribute{
			Key: "aria-label",
			Val: "Link issues",
		},
		client: *http.DefaultClient,
	}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *htmlParser) Find() ([]string, error) {
	prPage, err := p.downloadPRPage()
	if err != nil {
		return nil, err
	}

	nodes, err := p.parseHTMLNodes(prPage)
	if err != nil {
		return nil, err
	}

	tagNode := p.findTagNodeWithIssueLinks(nodes)
	if tagNode == nil {
		return nil, fmt.Errorf("tag %q with attribute %s=%q does not exist", p.tag, p.attr.Key, p.attr.Val)
	}

	return findIssueLinks(tagNode)
}

func (p *htmlParser) downloadPRPage() ([]byte, error) {
	if p.prURL == "" {
		return nil, fmt.Errorf("PR URL must not be empty")
	}
	response, err := p.client.Get(p.prURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func (p *htmlParser) parseHTMLNodes(page []byte) (*html.Node, error) {
	return html.Parse(strings.NewReader(string(page)))
}

func (p *htmlParser) findTagNodeWithIssueLinks(n *html.Node) *html.Node {
	// If the current node is the desired tag, return it.
	if n.Type == html.ElementNode && n.Data == p.tag {
		for _, a := range n.Attr {
			if a.Key == p.attr.Key && a.Val == p.attr.Val {
				return n
			}
		}
	}
	// Recursively look into the next nodes for the desired tag.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := p.findTagNodeWithIssueLinks(c)
		if node != nil {
			return node
		}
	}
	return nil
}

func findIssueLinks(tag *html.Node) ([]string, error) {
	links := make([]string, 0)

	var traverse func(*html.Node)

	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && strings.Contains(a.Val, "/issues/") {
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(tag)

	if len(links) == 0 {
		return nil, fmt.Errorf("no links found inside the provided tag: %s", tag.Data)
	}
	return links, nil
}

func SetAttribute(key, val string) func(p *htmlParser) {
	return func(p *htmlParser) {
		if key != "" {
			p.attr.Key = key
		}
		if val != "" {
			p.attr.Val = val
		}
	}
}

func SetTag(tag string) func(p *htmlParser) {
	return func(p *htmlParser) {
		if tag != "" {
			p.tag = tag
		}
	}
}
