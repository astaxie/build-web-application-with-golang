// As of Google go 1.1.2, `println()` and `print()` are hidden functions included from the runtime package.
// However it's encouraged to use the print functions from the `fmt` package.
package main

import "fmt"

func f() {
	fmt.Println("First")
	print("Second ")
	println(" Third")
}
func main() {
	f()
}
