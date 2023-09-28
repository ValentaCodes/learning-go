package main

import (
	"fmt"

	"example.com/greetings"
	"golang.org/x/example/hello/reverse"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Bobby")
	fmt.Println(message)
	reverseStringAndInt()
}

func reverseStringAndInt() {
	fmt.Println(reverse.String("Hello World"), reverse.Int(2345), reverse.Float(23234.32))
}
