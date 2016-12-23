// Código de exemplo do capítulo 2.3 de "Build Web Application with Golang"
// Propósito: mostrar como usar `panic()` e `recover()`
package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func check_user() {
	if user == "" {
		panic("no value for $USER")
	}
	fmt.Println("Environment Variable `USER` =", user)
}
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Panic message =", x);
			b = true
		}
	}()
	f() // se f causar pânico, isto irá recuperar
	return
}
func main(){
	didPanic := throwsPanic(check_user)
	fmt.Println("didPanic =", didPanic)
}
