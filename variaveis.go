package main

import (
	"fmt"
)

var ( // Declarando multiplas variáveis
	nome       string = "Cesar"
	valor      int    = 42
	a1, a2, a3        = 10, 20, 30
)

const pi = 3.141592

func main() {
	var a string
	a = "texto 1"
	b := "texto 2"

	ola := func() {
		fmt.Printf("Olá da função anônima!\r\n")
	}

	fmt.Printf("a tipo: %T\r\n", a)
	fmt.Printf("b tipo: %T\r\n", b)
	fmt.Printf("π tipo: %T\r\n", pi)
	fmt.Printf("ola tipo: %T\r\n", ola)

	fmt.Printf("valor de a = %v\r\n", a)
	fmt.Printf("valor de b = %v\r\n", b)
	fmt.Printf("valor de π = %v\r\n", pi)

	ola()
}
