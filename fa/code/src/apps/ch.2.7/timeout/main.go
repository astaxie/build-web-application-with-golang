// Example code for Chapter 2.7 from "Build Web Application with Golang"
// Purpose: Shows how to create and use a timeout
package main

import (
	"fmt"
	"time"
)


func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
