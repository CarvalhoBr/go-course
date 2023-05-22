package main

import "fmt"

type endereco struct {
	rua    string
	numero uint16
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {

	var u usuario

	fmt.Println(u)

	u.nome = "Brandon"
	u.idade = 22

	fmt.Println(u)

	u2 := usuario{
		nome:  "Mirella",
		idade: 27,
		endereco: endereco{
			rua:    "Alberto Lamego",
			numero: 405,
		},
	}

	fmt.Println(u2)
}
