// Código de exemplo da seção 2.4 de "Build Web Application with Golang"
// Propósito: Outro exemplo de campos incorporados
package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // estrutura como campo incorporado
	Skills     // string slice como campo incorporado
	int        // tipo embutido como campo incorporado
	speciality string
}

func main() {
	// inicializa Student Jane
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// acessa os campos
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// modifica o valor do campo skill
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// modifica campo incorporado
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}
