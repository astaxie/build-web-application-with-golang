// Código de exemplo do capítulo 2.3 de "Build Web Application with Golang"
// Propósito: Criando uma função básica
package main

import "fmt"

// retorna o maior valor entre a e b
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

	max_xy := max(x, y) // chama a função max(x, y)
	max_xz := max(x, z) // chama a função max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // chama a função aqui
}
