// Example code for Chapter 2.2 from "Build Web Application with Golang"
// Purpose: Goes over the assignment and manipulation of basic data types.
package main

import (
	"errors"
	"fmt"
)

// constants
const Pi = 3.1415926

// booleans default to `false`
var isActive bool                   // global variable
var enabled, disabled = true, false // omit type of variables

// grouped definitions
const (
	i         = 1e4
	MaxThread = 10
	prefix    = "astaxie_"
)

var (
	frenchHello string      // basic form to define string
	emptyString string = "" // define a string with empty string
)

func show_multiple_assignments() {
	fmt.Println("show_multiple_assignments()")
	var v1 int = 42

	// Define three variables with type "int", and initialize their values.
	// vname1 is v1, vname2 is v2, vname3 is v3
	var v2, v3 int = 2, 3

	// `:=` only works in functions
	// `:=` is the short way of declaring variables without
	//  specifying the type and using the keyboard `var`.
	vname1, vname2, vname3 := v1, v2, v3

	// `_` disregards the returned value.
	_, b := 34, 35

	fmt.Printf("vname1 = %v, vname2 = %v, vname3 = %v\n", vname1, vname2, vname3)
	fmt.Printf("v1 = %v, v2 = %v, v3 = %v\n", v1, v2, v3)
	fmt.Println("b =", b)
}
func show_bool() {
	fmt.Println("show_bool()")
	var available bool // local variable
	valid := false     // Shorthand assignment
	available = true   // assign value to variable

	fmt.Printf("valid = %v, !valid = %v\n", valid, !valid)
	fmt.Printf("available = %v\n", available)
}
func show_different_types() {
	fmt.Println("show_different_types()")
	var (
		unicodeChar rune
		a           int8
		b           int16
		c           int32
		d           int64
		e           byte
		f           uint8
		g           int16
		h           uint32
		i           uint64
	)
	var cmplx complex64 = 5 + 5i

	fmt.Println("Default values for int types")
	fmt.Println(unicodeChar, a, b, c, d, e, f, g, h, i)

	fmt.Printf("Value is: %v\n", cmplx)
}
func show_strings() {
	fmt.Println("show_strings()")
	no, yes, maybe := "no", "yes", "maybe" // brief statement
	japaneseHello := "Ohaiyou"
	frenchHello = "Bonjour" // basic form of assign values

	fmt.Println("Random strings")
	fmt.Println(frenchHello, japaneseHello, no, yes, maybe)

	// The backtick, `, will not escape any character in a string
	fmt.Println(`This 
	is on
	multiple lines`)
}
func show_string_manipulation() {
	fmt.Println("show_string_manipulation()")
	var s string = "hello"

	//You can't do this with strings
	//s[0] = 'c'

	s = "hello"
	c := []byte(s) // convert string to []byte type
	c[0] = 'c'
	s2 := string(c) // convert back to string type

	m := " world"
	a := s + m

	d := "c" + s[1:] // you cannot change string values by index, but you can get values instead.
	fmt.Printf("%s\n", d)

	fmt.Printf("s = %s, c = %v\n", s, c)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("combined strings\na = %s, d = %s\n", a, d)
}
func show_errors() {
	fmt.Println("show_errors()")
	err := errors.New("Example error message\n")
	if err != nil {
		fmt.Print(err)
	}
}
func show_iota() {
	fmt.Println("show_iota()")
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // If there is no expression after constants name,
		// it uses the last expression, so here is saying w = iota implicitly.
		// Therefore w == 3, and y and x both can omit "= iota" as well.
	)

	const v = iota // once iota meets keyword `const`, it resets to `0`, so v = 0.

	const (
		e, f, g = iota, iota, iota // e=0,f=0,g=0 values of iota are same in one line.
	)
	fmt.Printf("x = %v, y = %v, z = %v, w = %v\n", x, y, z, w)
	fmt.Printf("v = %v\n", v)
	fmt.Printf("e = %v, f = %v, g = %v\n", e, f, g)
}

// Functions and variables starting with a capital letter are public to other packages.
// Everything else is private.
func This_is_public()  {}
func this_is_private() {}

func set_default_values() {
	// default values for the types.
	const (
		a int     = 0
		b int8    = 0
		c int32   = 0
		d int64   = 0
		e uint    = 0x0
		f rune    = 0   // the actual type of rune is int32
		g byte    = 0x0 // the actual type of byte is uint8
		h float32 = 0   // length is 4 byte
		i float64 = 0   //length is 8 byte
		j bool    = false
		k string  = ""
	)
}
func show_arrays() {
	fmt.Println("show_arrays()")
	var arr [10]int // an array of type int
	arr[0] = 42     // array is 0-based
	arr[1] = 13     // assign value to element

	a := [3]int{1, 2, 3} // define a int array with 3 elements

	b := [10]int{1, 2, 3}
	// define a int array with 10 elements,
	// and first three are assigned, rest of them use default value 0.

	c := [...]int{4, 5, 6} // use `…` replace with number of length, Go will calculate it for you.

	// define a two-dimensional array with 2 elements, and each element has 4 elements.
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// You can write about declaration in a shorter way.
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println("arr =", arr)
	fmt.Printf("The first element is %d\n", arr[0]) // get element value, it returns 42
	fmt.Printf("The last element is %d\n", arr[9])
	//it returns default value of 10th element in this array, which is 0 in this case.

	fmt.Println("array a =", a)
	fmt.Println("array b =", b)
	fmt.Println("array c =", c)

	fmt.Println("array doubleArray =", doubleArray)
	fmt.Println("array easyArray =", easyArray)
}
func show_slices() {
	fmt.Println("show_slices()")
	// define a slice with 10 elements which types are byte
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// define two slices with type []byte
	var a, b []byte

	// a points to elements from 3rd to 5th in array ar.
	a = ar[2:5]
	// now a has elements ar[2]、ar[3] and ar[4]

	// b is another slice of array ar
	b = ar[3:5]
	// now b has elements ar[3] and ar[4]

	// define an array
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// define two slices
	var aSlice, bSlice []byte

	// some convenient operations
	aSlice = array[:3] // equals to aSlice = array[0:3] aSlice has elements a,b,c
	aSlice = array[5:] // equals to aSlice = array[5:10] aSlice has elements f,g,h,i,j
	aSlice = array[:]  // equals to aSlice = array[0:10] aSlice has all elements

	// slice from slice
	aSlice = array[3:7]  // aSlice has elements d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice contains aSlice[1], aSlice[2], so it has elements e,f
	bSlice = aSlice[:3]  // bSlice contains aSlice[0], aSlice[1], aSlice[2], so it has d,e,f
	bSlice = aSlice[0:5] // slcie could be expanded in range of cap, now bSlice contains d,e,f,g,h
	bSlice = aSlice[:]   // bSlice has same elements as aSlice does, which are d,e,f,g

	fmt.Println("slice ar =", ar)
	fmt.Println("slice a =", a)
	fmt.Println("slice b =", b)
	fmt.Println("array =", array)
	fmt.Println("slice aSlice =", aSlice)
	fmt.Println("slice bSlice =", bSlice)
	fmt.Println("len(bSlice) =", len(bSlice))
}
func show_map() {
	fmt.Println("show_map()")
	// use string as key type, int as value type, and you have to use `make` initialize it.
	var numbers map[string]int
	// another way to define map
	numbers = make(map[string]int)
	numbers["one"] = 1 // assign value by key
	numbers["ten"] = 10
	numbers["three"] = 3

	// Initialize a map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

	fmt.Println("map numbers =", numbers)
	fmt.Println("The third number is: ", numbers["three"]) // get values
	// It prints: The third number is: 3

	// map has two return values. For second value, if the key doesn't exist，ok is false，true otherwise.
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // delete element with key "c"
	fmt.Printf("map rating = %#v\n", rating)
}
func main() {
	show_multiple_assignments()
	show_bool()
	show_different_types()
	show_strings()
	show_string_manipulation()
	show_errors()
	show_iota()
	set_default_values()
	show_arrays()
	show_slices()
	show_map()
}
