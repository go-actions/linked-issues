package io

import (
	"fmt"
	"strings"
)

type Output struct {
	Name  string
	Value Formatter
}

func Print(o []Output) {
	for i := range o {
		fmt.Printf("::set-output name=%s::%s\n", o[i].Name, o[i].Value.Format())
	}
}

const (
	IssueNumber      = "IssueNumber"
	IssueURL         = "IssueURL"
	ExternalIssueRef = "ExternalIssueRef"
)

type Formatter interface {
	Format() string
}

func NewFormatter(issues []string, format string) Formatter {
	switch format {
	case IssueNumber:
		return &numberFormatter{issues: issues}
	case IssueURL:
		return &urlFormatter{issues: issues}
	case ExternalIssueRef:
		return &externalIssueRefFormatter{issues: issues}
	default:
		return &numberFormatter{issues: issues} // By default, use numberFormatter.
	}
}

type numberFormatter struct {
	issues []string
}

func (nf *numberFormatter) Format() string {
	output := make([]string, 0)

	for i := range nf.issues {
		parts := strings.Split(nf.issues[i], "/")
		output = append(output, parts[len(parts)-1])
	}
	return strings.Join(output, " ")
}

type urlFormatter struct {
	issues []string
}

func (uf *urlFormatter) Format() string {
	return strings.Join(uf.issues, " ")
}

type externalIssueRefFormatter struct {
	issues []string
}

func (ref *externalIssueRefFormatter) Format() string {
	output := make([]string, 0)

	for i := range ref.issues {
		parts := strings.Split(ref.issues[i], "/")
		output = append(output, fmt.Sprintf("%s/%s#%s", parts[len(parts)-4], parts[len(parts)-3], parts[len(parts)-1]))
	}
	return strings.Join(output, " ")
}
