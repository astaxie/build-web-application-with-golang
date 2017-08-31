// Código de exemplo para o Capítulo 1.2 do "Build Web Application with Golang"
// Propósito: Execute este arquivo para verificar se o seu workspace está corretamente configurado.
// Para executar, navegue até o diretório onde ele estiver salvo e digite no console `go run main.go`
// Se o texto "Hello World" não aparecer, então configure seu ambiente novamente.
package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
