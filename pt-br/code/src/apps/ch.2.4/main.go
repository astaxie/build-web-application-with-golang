// Código de exemplo da seção 2.4 de "Build Web Application with Golang"
// Propósito: Mostrar formas diferentes de criar uma estrutura
package main

import "fmt"

func show_basic_struct() {
	fmt.Println("\nshow_basic_struct()")
	type person struct {
		name string
		age  int
	}

	var P person // p é do tipo person

	P.name = "Astaxie"                              // atribui "Astaxie" ao campo 'name' de p
	P.age = 25                                      // atribui 25 ao campo 'age' de p
	fmt.Printf("The person's name is %s\n", P.name) // acessa o campo 'name' de p

	tom := person{"Tom", 25}

	bob := person{age: 24, name: "Bob"}

	fmt.Printf("tom = %+v\n", tom)
	fmt.Printf("bob = %#v\n", bob)
}
func show_anonymous_struct() {
	fmt.Println("\nshow_anonymous_struct()")
	fmt.Printf("Anonymous struct = %#v\n", struct {
		name  string
		count int
	}{
		"counter", 1,
	})
}
func main() {
	show_basic_struct()
	show_anonymous_struct()
}
