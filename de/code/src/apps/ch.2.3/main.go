// Example code for Chapter 2.3 from "Build Web Application with Golang"
// Purpose: Goes over if, else, switch conditions, loops and defer.
package main

import "fmt"

func computedValue() int {
	return 1
}
func show_if() {
	fmt.Println("\n#show_if()")
	x := computedValue()
	integer := 23

	fmt.Println("x =", x)
	fmt.Println("integer =", integer)
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	if integer == 3 {
		fmt.Println("The integer is equal to 3")
	} else if integer < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}
}
func show_if_var() {
	fmt.Println("\n#show_if_var()")
	// initialize x, then check if x greater than
	if x := computedValue(); x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	// the following code will not compile, since `x` is only accessible with the if/else block
	// fmt.Println(x)
}
func show_goto() {
	fmt.Println("\n#show_goto()")
	// The call to the label switches the goroutine it seems.
	i := 0
Here: // label ends with ":"
	fmt.Println(i)
	i++
	if i < 10 {
		goto Here // jump to label "Here"
	}
}
func show_for_loop() {
	fmt.Println("\n#show_for_loop()")
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("part 1, sum is equal to ", sum)

	sum = 1
	// The compiler will remove the `;` from the line below.
	// for ; sum < 1000 ; {
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("part 2, sum is equal to ", sum)

	for index := 10; 0 < index; index-- {
		if index == 5 {
			break // or continue
		}
		fmt.Println(index)
	}

}
func show_loop_through_map() {
	fmt.Println("\n#show_loop_through_map()")
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("map value = ", m)
	for k, v := range m {
		fmt.Println("map's key: ", k)
		fmt.Println("map's value: ", v)
	}
}
func show_switch() {
	fmt.Println("\n#show_switch()")
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	integer := 6
	fmt.Println("integer =", integer)
	switch integer {
	case 4:
		fmt.Println("integer == 4")
		fallthrough
	case 5:
		fmt.Println("integer <= 5")
		fallthrough
	case 6:
		fmt.Println("integer <= 6")
		fallthrough
	case 7:
		fmt.Println("integer <= 7")
		fallthrough
	case 8:
		fmt.Println("integer <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
func show_defer() {
	fmt.Println("\nshow_defer()")
	defer fmt.Println("(last defer)")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
func main() {
	show_if()
	show_if_var()
	show_goto()
	show_for_loop()
	show_loop_through_map()
	show_switch()
	show_defer()
}
