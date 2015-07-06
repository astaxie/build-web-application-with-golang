// Example code for Chapter 2.4 from "Build Web Application with Golang"
// Purpose: Shows a name conflict with a embedded field
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string // Human has phone field
}

type Employee struct {
	Human      // embedded field Human
	speciality string
	phone      string // phone in employee
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// access phone field in Human
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
