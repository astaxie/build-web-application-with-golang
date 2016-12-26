// Código de exemplo do capítulo 2.3 de "Build Web Application with Golang"
// Propósito: mostra como definir um tipo de função
package main

import "fmt"

type testInt func(int) bool // define uma função como um tipo de variável

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// passa a função `f` como um argumento para outra função

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
func init() {
	fmt.Println("\n#init() was called.")
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // usa funções como valores
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}
