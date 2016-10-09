// Example code for Chapter 1.2 from "Build Web Application with Golang"
// Purpose: Shows how to create a simple package called `mymath`
// This package must be imported from another go file to run.
package mymath

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
