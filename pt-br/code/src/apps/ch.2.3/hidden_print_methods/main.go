// A partir do Google go 1.1.2, `println()` e `print()` são funções ocultas incluídas no pacote de tempo de execução.
// No entanto, é encorajado utilizar as funções de impressão do pacote `fmt`
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
