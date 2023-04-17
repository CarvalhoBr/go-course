package main

import "fmt"

func main() {
	// Tipagem explicita
	var variavel1 string = "Variavel 1"

	// Tipagem implicita
	variavel2 := "Variavel 2"

	// Declaração multipla
	var (
		variavel3 string = "test"
		variavel4 string = "test"
	)

	// Declaração multipla com inferencia de tipos
	variavel5, variavel6 := "Variavel5", 2

	// Constante
	const constante1 = 1

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	fmt.Println(constante1)
}
