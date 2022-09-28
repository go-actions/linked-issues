package io

import (
	"fmt"
	"strings"
)

type Output struct {
	Name   string
	Value  Formatter
	Issues []string
}

func Print(o []Output) {
	for i := range o {
		fmt.Printf("::set-output name=%s::%s\n", o[i].Name, o[i].Value.Format(o[i].Issues))
	}
}

const (
	IssueNumber      = "IssueNumber"
	IssueURL         = "IssueURL"
	ExternalIssueRef = "ExternalIssueRef"
)

type Formatter interface {
	Format(issues []string) string
}

func NewFormatter(format string) Formatter {
	switch format {
	case IssueNumber:
		return &numberFormatter{}
	case IssueURL:
		return &urlFormatter{}
	case ExternalIssueRef:
		return &externalIssueRefFormatter{}
	default:
		return &numberFormatter{} // By default, use numberFormatter.
	}
}

type numberFormatter struct {
}

func (nf *numberFormatter) Format(issues []string) string {
	output := make([]string, 0)

	for i := range issues {
		parts := strings.Split(issues[i], "/")
		output = append(output, parts[len(parts)-1])
	}
	return strings.Join(output, " ")
}

type urlFormatter struct {
}

func (uf *urlFormatter) Format(issues []string) string {
	return strings.Join(issues, " ")
}

type externalIssueRefFormatter struct {
}

func (ref *externalIssueRefFormatter) Format(issues []string) string {
	output := make([]string, 0)

	for i := range issues {
		parts := strings.Split(issues[i], "/")
		output = append(output, fmt.Sprintf("%s/%s#%s", parts[len(parts)-4], parts[len(parts)-3], parts[len(parts)-1]))
	}
	return strings.Join(output, " ")
}
