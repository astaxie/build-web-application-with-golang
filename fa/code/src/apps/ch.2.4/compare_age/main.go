// Example code for Chapter 2.4 from "Build Web Application with Golang"
// Purpose: Shows you how to pass and use structs.
package main

import "fmt"

// define a new type
type person struct {
	name string
	age  int
}

// compare age of two people, return the older person and differences of age
// struct is passed by value
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person

	// initialization
	tom.name, tom.age = "Tom", 18

	// initialize two values by format "field:value"
	bob := person{age: 25, name: "Bob"}

	// initialize two values with order
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)
}
