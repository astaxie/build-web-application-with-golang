// Example code for Chapter 1.2 from "Build Web Application with Golang"
// Purpose: Run this file to check if your workspace is setup correctly.
// To run, navigate to the current directory in a console and type `go run main.go`
// If the text "Hello World" isn't shown, then setup your workspace again.
package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
