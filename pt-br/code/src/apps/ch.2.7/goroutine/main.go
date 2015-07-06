// Example code for Chapter 2.7 from "Build Web Application with Golang"
// Purpose: Shows how to launch a simple gorountine
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") // create a new goroutine
	say("hello")    // current goroutine
}
