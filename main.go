package main

import (
	"fmt"
	"os"

	"github.com/hossainemruz/linked-issues/finder"
	"github.com/hossainemruz/linked-issues/io"
)

func main() {
	input := io.ReadInput()

	f := finder.NewIssueFinder(
		finder.NewHTMLParser(
			input.GetPRUrl(),
			finder.SetAttribute(input.GetAttrKey(), input.GetAttrVal()),
			finder.SetTag(input.GetTag()),
		),
	)

	issues, err := f.Find()

	if err != nil {
		fmt.Println("Failed to find the linked Issues. Reason: ", err.Error())
		os.Exit(1)
	}

	output := []io.Output{
		{
			Name:  "issues",
			Value: io.NewFormatter(issues, input.GetFormat()),
		},
	}
	io.Print(output)
}
