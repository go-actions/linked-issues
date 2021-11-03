package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var person string

func init() {
	flag.StringVar(&person, "person", "", "Person to greet.")
	flag.Parse()
}

func main() {
	args := os.Args
	fmt.Println("All arguments: ", args)
	fmt.Printf("Congratulation %s!!!. Your Github Action worked.\n", person)
	fmt.Printf("::set-output name=time::%s\n", time.Now().String())
}
