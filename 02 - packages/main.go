package main

import (
	"fmt"
	"packages/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("brandonnotemail.com")
	fmt.Println(erro)
}
