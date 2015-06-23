// Example code for Chapter 2.4 from "Build Web Application with Golang"
// Purpose: Shows different ways of creating a struct
package main

import "fmt"

func show_basic_struct() {
	fmt.Println("\nshow_basic_struct()")
	type person struct {
		name string
		age  int
	}

	var P person // p is person type

	P.name = "Astaxie"                              // assign "Astaxie" to the filed 'name' of p
	P.age = 25                                      // assign 25 to field 'age' of p
	fmt.Printf("The person's name is %s\n", P.name) // access field 'name' of p

	tom := person{"Tom", 25}

	bob := person{age: 24, name: "Bob"}

	fmt.Printf("tom = %+v\n", tom)
	fmt.Printf("bob = %#v\n", bob)
}
func show_anonymous_struct() {
	fmt.Println("\nshow_anonymous_struct()")
	fmt.Printf("Anonymous struct = %#v\n", struct {
		name  string
		count int
	}{
		"counter", 1,
	})
}
func main() {
	show_basic_struct()
	show_anonymous_struct()
}
