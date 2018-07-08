// Example code for Chapter 2.3 from "Build Web Application with Golang"
// Purpose: Showing how to use `panic()` and `recover()`
package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func check_user() {
	if user == "" {
		panic("no value for $USER")
	}
	fmt.Println("Environment Variable `USER` =", user)
}
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Panic message =", x);
			b = true
		}
	}()
	f() // if f causes panic, it will recover
	return
}
func main(){
	didPanic := throwsPanic(check_user)
	fmt.Println("didPanic =", didPanic)
}
