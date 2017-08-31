// Código de exemplo para o Capítulo 2.2 do "Build Web Application with Golang"
// Propósito: Revisar os conceitos de atribuição e manipulação de tipos de dados básicos.
package main

import (
	"errors"
	"fmt"
)

// constantes
const Pi = 3.1415926

// booleanos: padrão `false`
var isActive bool                   // variável global
var enabled, disabled = true, false // omite o tipo das variáveis

// definições agrupadas
const (
	i         = 1e4
	MaxThread = 10
	prefix    = "astaxie_"
)

var (
	frenchHello string      // forma básica para definir strings
	emptyString string = "" // define uma string vazia
)

func show_multiple_assignments() {
	fmt.Println("show_multiple_assignments()")
	var v1 int = 42

	// Define três variáveis com o tipo "int" e inicializa seus valores.
	// vname1 é v1, vname2 é v2, vname3 é v3
	var v2, v3 int = 2, 3

	// `:=` funciona somente em funções
	// `:=` é a maneira curta de declarar variáveis sem
	//  especificar o tipo e usando a palvra-chave `var`.
	vname1, vname2, vname3 := v1, v2, v3

	// `_` desconsidera o valor retornado.
	_, b := 34, 35

	fmt.Printf("vname1 = %v, vname2 = %v, vname3 = %v\n", vname1, vname2, vname3)
	fmt.Printf("v1 = %v, v2 = %v, v3 = %v\n", v1, v2, v3)
	fmt.Println("b =", b)
}
func show_bool() {
	fmt.Println("show_bool()")
	var available bool // variável local
	valid := false     // declara a variável e infere o tipo booleano (declaração curta)
	available = true   // atribui um valor a variável

	fmt.Printf("valid = %v, !valid = %v\n", valid, !valid)
	fmt.Printf("available = %v\n", available)
}
func show_different_types() {
	fmt.Println("show_different_types()")
	var (
		unicodeChar rune
		a           int8
		b           int16
		c           int32
		d           int64
		e           byte
		f           uint8
		g           int16
		h           uint32
		i           uint64
	)
	var cmplx complex64 = 5 + 5i

	fmt.Println("Default values for int types")
	fmt.Println(unicodeChar, a, b, c, d, e, f, g, h, i)

	fmt.Printf("Value is: %v\n", cmplx)
}
func show_strings() {
	fmt.Println("show_strings()")
	no, yes, maybe := "no", "yes", "maybe" // declaração curta
	japaneseHello := "Ohaiyou"
	frenchHello = "Bonjour" // forma básica de atribuir valores

	fmt.Println("Random strings")
	fmt.Println(frenchHello, japaneseHello, no, yes, maybe)

	// A crase, `, não deixa escapar qualquer caractere da string
	fmt.Println(`This 
	is on
	multiple lines`)
}
func show_string_manipulation() {
	fmt.Println("show_string_manipulation()")
	var s string = "hello"

	//Você não pode fazer isso com strings
	//s[0] = 'c'

	s = "hello"
	c := []byte(s) // converte string para o tipo []byte
	c[0] = 'c'
	s2 := string(c) // converte novamente para o tipo string

	m := " world"
	a := s + m

	d := "c" + s[1:] // você não pode alterar valores de string pelo índice, mas você pode obter os valores desta maneira
	fmt.Printf("%s\n", d)

	fmt.Printf("s = %s, c = %v\n", s, c)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("combined strings\na = %s, d = %s\n", a, d)
}
func show_errors() {
	fmt.Println("show_errors()")
	err := errors.New("Example error message\n")
	if err != nil {
		fmt.Print(err)
	}
}
func show_iota() {
	fmt.Println("show_iota()")
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // Se não houver nenhuma expressão após o nome da constante,
		// ela usa a última expressão, ou seja, neste caso w = iota implicitamente.
		// Portanto w == 3, e tanto y quanto z podem omitir "= iota" também.
	)

	const v = iota // uma vez que iota encontra a palavra-chave `const`, ela é redefinida para `0`, então v = 0.

	const (
		e, f, g = iota, iota, iota // e=0,f=0,g=0 valores de iota são iguais quando estão na mesma linha.
	)
	fmt.Printf("x = %v, y = %v, z = %v, w = %v\n", x, y, z, w)
	fmt.Printf("v = %v\n", v)
	fmt.Printf("e = %v, f = %v, g = %v\n", e, f, g)
}

// Funções e variáveis começando com uma letra maiúscula são públicas para outros pacotes.
// Todo o resto é privado.
func This_is_public()  {}
func this_is_private() {}

func set_default_values() {
	// valores padrão para os tipos.
	const (
		a int     = 0
		b int8    = 0
		c int32   = 0
		d int64   = 0
		e uint    = 0x0
		f rune    = 0   // o tipo real de rune é int32
		g byte    = 0x0 // o tipo real de byte é uint8
		h float32 = 0   // comprimento é 4 bytes
		i float64 = 0   // comprimento é 8 bytes
		j bool    = false
		k string  = ""
	)
}
func show_arrays() {
	fmt.Println("show_arrays()")
	var arr [10]int // um array do tipo int
	arr[0] = 42     // array é baseado em 0
	arr[1] = 13     // atribui valor ao elemento

	a := [3]int{1, 2, 3} // define um array do tipo int com 3 elementos

	b := [10]int{1, 2, 3}
	// define um array do tipo int com 10 elementos,
	// e os três primeiros são atribuídos, o restante usa o valor padrão 0.

	c := [...]int{4, 5, 6} // usa `…` para substituir o número do comprimento, Go irá calcular isto para você.

	// define um array bidimensional com 2 elementos, e onde cada elemento possui 4 elementos.
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// Você pode escrever a declaração de forma curta.
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println("arr =", arr)
	fmt.Printf("The first element is %d\n", arr[0]) // obtém o valor do elemento, retorna 42
	fmt.Printf("The last element is %d\n", arr[9])
	// retorna o valor padrão do décimo elemento do array, que neste caso é 0.

	fmt.Println("array a =", a)
	fmt.Println("array b =", b)
	fmt.Println("array c =", c)

	fmt.Println("array doubleArray =", doubleArray)
	fmt.Println("array easyArray =", easyArray)
}
func show_slices() {
	fmt.Println("show_slices()")
	// define um slice de tipo byte com 10 elementos
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// define dois slices com o tipo []byte
	var a, b []byte

	// a aponta para os elementos da posição 3 até a 5 do array ar.
	a = ar[2:5]
	// agora a possui os elementos ar[2], ar[3] e ar[4]

	// b é outro slice do array ar
	b = ar[3:5]
	// agora b possui os elementos de ar[3] e ar[4]

	// define um array
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// define dois slices
	var aSlice, bSlice []byte

	// algumas operações convenientes
	aSlice = array[:3] // igual a aSlice = array[0:3] aSlice possui os elementos a,b,c
	aSlice = array[5:] // igual a aSlice = array[5:10] aSlice possui os elementos f,g,h,i,j
	aSlice = array[:]  // igual a aSlice = array[0:10] aSlice possui todos os elementos

	// slice de slice
	aSlice = array[3:7]  // aSlice possui os elementos d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice contém aSlice[1], aSlice[2], ou seja, ele possui os elementos e,f
	bSlice = aSlice[:3]  // bSlice contém aSlice[0], aSlice[1], aSlice[2], ou seja, ele possui os elementos d,e,f
	bSlice = aSlice[0:5] // bSlice pode ser expandido de acordo com cap, agora bSlice contém d,e,f,g,h
	bSlice = aSlice[:]   // bSlice possui os mesmos elementos de aSlice, os quais são d,e,f,g

	fmt.Println("slice ar =", ar)
	fmt.Println("slice a =", a)
	fmt.Println("slice b =", b)
	fmt.Println("array =", array)
	fmt.Println("slice aSlice =", aSlice)
	fmt.Println("slice bSlice =", bSlice)
	fmt.Println("len(bSlice) =", len(bSlice))
}
func show_map() {
	fmt.Println("show_map()")
	// use tipo string para a chave, tipo int para o valor, e você precisa usar `make` para inicializar
	var numbers map[string]int
	// outra forma de declarar map
	numbers = make(map[string]int)
	numbers["one"] = 1 // atribui valor pela chave
	numbers["ten"] = 10
	numbers["three"] = 3

	// Inicializa um map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

	fmt.Println("map numbers =", numbers)
	fmt.Println("The third number is: ", numbers["three"]) // obtém os valores
	// Isto irá imprimir: The third number is: 3

	// map possui dois valores de retorno. Se a chave não existir, o segundo valor, neste caso "ok", será falso (false). Caso contrário será verdadeiro (true).
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // deleta o elemento com a chave "c"
	fmt.Printf("map rating = %#v\n", rating)
}
func main() {
	show_multiple_assignments()
	show_bool()
	show_different_types()
	show_strings()
	show_string_manipulation()
	show_errors()
	show_iota()
	set_default_values()
	show_arrays()
	show_slices()
	show_map()
}
