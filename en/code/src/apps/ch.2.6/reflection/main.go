package main

import (
	"fmt"
	"reflect"
)

func show_interface_none() {
	fmt.Println("\nshow_interface_none()")
	var a interface{}
	a = "string"
	a = 1
	a = false
	fmt.Println("a =", a)
}
func show_reflection() {
	fmt.Println("\nshow_reflection()")
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	
	p := reflect.ValueOf(&x)
	newX := p.Elem()
	newX.SetFloat(7.1)
	fmt.Println("newX =", newX)
	fmt.Println("newX float64() value:", newX.Float())
}
func main() {
	show_interface_none()
	show_reflection()
}
