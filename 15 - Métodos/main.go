package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando")
}

func (u usuario) eMaiordeIdade() bool {
	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()

	fmt.Printf("Usuário é maior de idade: %v", usuario1.eMaiordeIdade())
}
