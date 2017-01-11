// Código de exemplo da seção 2.4 de "Build Web Application with Golang"
// Propósito: Mostra um conflito de nomes com um campo incorporado
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string // Human possui campo phone
}

type Employee struct {
	Human      // campo incorporado de Human
	speciality string
	phone      string // phone em employee
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// acessa o campo phone em Human
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
