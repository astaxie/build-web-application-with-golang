// Example code for Chapter 2.3 from "Build Web Application with Golang"
// Purpose: Shows how to return multiple values from a function
package main

import "fmt"

// return results of A + B and A * B
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
