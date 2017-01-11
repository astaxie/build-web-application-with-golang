// Código de exemplo da seção 2.4 de "Build Web Application with Golang"
// Propósito: Mostrar como utilizar estruturas em Go.
package main

import "fmt"

// define um novo tipo
type person struct {
	name string
	age  int
}

// compara a idade de duas pessoas e retorna dois valores: a pessoa mais velha e a diferença de idade
// estrutura é passada por valor
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person

	// inicialização
	tom.name, tom.age = "Tom", 18

	// inicializa dois valores pelo formato "campo:valor"
	bob := person{age: 25, name: "Bob"}

	// inicializa dois valores em ordem
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)
}
