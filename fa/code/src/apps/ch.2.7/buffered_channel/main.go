// Example code for Chapter 2.7 from "Build Web Application with Golang"
// Purpose: Shows how to use a buffered channel
package main

import "fmt"

func main() {
	c := make(chan int, 2) // change 2 to 1 will have runtime error, but 3 is fine
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
