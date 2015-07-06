// Example code for Chapter 2.4 from "Build Web Application with Golang"
// Purpose: Example of embedded fields
package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // anonymous field, it means Student struct includes all fields that Human has.
	speciality string
}

func main() {
	// initialize a student
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	// access fields
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// modify notes
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// modify age
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// modify weight
	fmt.Println("Mark is not an athlete any more")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}
