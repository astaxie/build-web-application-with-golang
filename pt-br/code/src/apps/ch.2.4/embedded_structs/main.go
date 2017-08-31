// Código de exemplo da seção 2.4 de "Build Web Application with Golang"
// Propósito: Exemplo de campos incorporados
package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // campo anônimo, significa que a estrutura Student inclui todos os campos que Human possui.
	speciality string
}

func main() {
	// inicializa um estudante
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	// acessa os campos
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// modifica especialidade
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// modifica idade
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// modifica peso
	fmt.Println("Mark is not an athlete any more")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}
