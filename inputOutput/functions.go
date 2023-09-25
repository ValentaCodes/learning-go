package main

// NOTE: this is a safe way to print
import "fmt"

func printData() {
	// print and printLN are only useful for development purposes
	// these functions are not guarenteed to work on every os and platform 
	println(name)

	fmt.Print("Hello")
}

