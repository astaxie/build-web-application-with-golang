// Example code for Chapter 2.3 from "Build Web Application with Golang"
// Purpose: Creating a basic function
package main

import "fmt"

// return greater value between a and b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) // call function max(x, y)
	max_xz := max(x, z) // call function max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // call function here
}
