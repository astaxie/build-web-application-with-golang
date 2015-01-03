// Código de ejemplo para el capítulo 1.2 del libro "Construye Aplicaciones Web con Golang"
// Propósito: Ejecuta este archivo para verificar que tu espacio de trabajo está configurado correctamente.
// Para ejecutarlo, busca el directorio actual en la terminal y escribe `go run main.go`
// Si el texto "Hello World" no se muestra, entonces configura tu ambiente de trabajo nuevamente.
package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
