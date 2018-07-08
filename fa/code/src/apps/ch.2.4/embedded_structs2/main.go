// Example code for Chapter 2.4 from "Build Web Application with Golang"
// Purpose: Another example of embedded fields
package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // struct as embedded field
	Skills     // string slice as embedded field
	int        // built-in type as embedded field
	speciality string
}

func main() {
	// initialize Student Jane
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// access fields
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// modify value of skill field
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// modify embedded field
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}
