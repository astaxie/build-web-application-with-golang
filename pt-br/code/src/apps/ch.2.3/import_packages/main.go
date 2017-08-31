// Código de exemplo do capítulo 2.3 de "Build Web Application with Golang"
// Propósito: mostra diferentes formas de importar um pacote.
// Nota: para o pacote `only_call_init`, fazemos referência ao caminho a partir do diretório
// base de `$GOPATH/src`. Golang desencoraja o uso de caminhos relativos para importar pacotes. 
// BAD: "./only_call_init"
// GOOD: "apps/ch.2.3/import_packages/only_call_init"
package main

import (
	// `_` irá chamar apenas init() dentro do pacote only_call_init
	_ "apps/ch.2.3/import_packages/only_call_init"
	f "fmt"         // importa o pacote como `f`
	. "math"        // torna os métodos públicos e constantes globais
	"mymath"        // pacote personalizado localizado em $GOPATH/src/
	"os"            // import normal de um pacote padrão
	"text/template" // o pacote leva o nome do último caminho da pasta, `template`
)

func main() {
	f.Println("mymath.Sqrt(4) =", mymath.Sqrt(4))
	f.Println("E =", E) // referencia math.E

	t, _ := template.New("test").Parse("Pi^2 = {{.}}")
	t.Execute(os.Stdout, Pow(Pi, 2))
}
