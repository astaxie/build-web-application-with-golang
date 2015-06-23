// Example code for Chapter 2.3 from "Build Web Application with Golang"
// Purpose: Shows different ways of importing a package.
// Note: For the package `only_call_init`, we reference the path from the
// base directory of `$GOPATH/src`. The reason being Golang discourage
// the use of relative paths when import packages.
// BAD: "./only_call_init"
// GOOD: "apps/ch.2.3/import_packages/only_call_init"
package main

import (
	// `_` will only call init() inside the package only_call_init
	_ "apps/ch.2.3/import_packages/only_call_init"
	f "fmt"         // import the package as `f`
	. "math"        // makes the public methods and constants global
	"mymath"        // custom package located at $GOPATH/src/
	"os"            // normal import of a standard package
	"text/template" // the package takes the name of last folder path, `template`
)

func main() {
	f.Println("mymath.Sqrt(4) =", mymath.Sqrt(4))
	f.Println("E =", E) // references math.E

	t, _ := template.New("test").Parse("Pi^2 = {{.}}")
	t.Execute(os.Stdout, Pow(Pi, 2))
}
