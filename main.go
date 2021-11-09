package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hossainemruz/linked-issues/io"
	"github.com/hossainemruz/linked-issues/parser"
	"golang.org/x/net/html"
)

type Input struct {
	prURL  string
	format string
	tag    string
	attr   html.Attribute
}

var input Input

func init() {
	flag.StringVar(&input.prURL, "pr-url", "", "URL of the pull request.")
	flag.StringVar(&input.tag, "tag", "form", "HTML tag which contans the issue links.")
	flag.StringVar(&input.attr.Key, "attr-key", "aria-label", "The key of the tag attribute.")
	flag.StringVar(&input.attr.Val, "attr-val", "Link issues", "The value of the tag attribute.")
	flag.StringVar(&input.format, "format", io.IssueNumber, "Format of the output issue list.")
}

func main() {
	// Parse the flags
	flag.Parse()

	//  Find the linked issues
	p := parser.NewHTMLParser()
	issues, err := p.Parse()
	if err != nil {
		fmt.Println("Failed to find the linked Issues. Reason: ", err.Error())
		os.Exit(1)
	}

	output := []io.Output{
		{
			Name:  "issues",
			Value: io.NewFormatter(issues, input.format),
		},
	}
	io.Print(output)
}
