package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	person := os.Getenv("PERSON")
	fmt.Printf("Congratulation %s!!!. Your Github Action worked.", person)
	fmt.Printf("::set-output name=time::%s", time.Now().String())
}
