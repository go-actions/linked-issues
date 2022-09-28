package io

import (
	"flag"
)

type Input struct {
	prURL   string
	format  string
	tag     string
	attrKey string
	attrVal string
}

func ReadInput() Input {
	var input Input
	flag.StringVar(&input.prURL, "pr-url", input.prURL, "URL of the pull request.")
	flag.StringVar(&input.tag, "tag", input.tag, "HTML tag which contans the issue links.")
	flag.StringVar(&input.attrKey, "attr-key", input.attrKey, "The key of the tag attribute.")
	flag.StringVar(&input.attrVal, "attr-val", input.attrVal, "The value of the tag attribute.")
	flag.StringVar(&input.format, "format", IssueNumber, "Format of the output issue list.")
	flag.Parse()

	return input
}

func (i *Input) GetFormat() string {
	return i.format
}

func (i *Input) GetPRUrl() string {
	return i.prURL
}

func (i *Input) GetTag() string {
	return i.tag
}

func (i *Input) GetAttrKey() string {
	return i.attrKey
}

func (i *Input) GetAttrVal() string {
	return i.attrVal
}
