// Código de exemplo do capítulo 2.3 de "Build Web Application with Golang"
// Propósito: mostra como retornar múltiplos valores de uma função
package main

import "fmt"

// retorna os resultados de A + B e A * B
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
