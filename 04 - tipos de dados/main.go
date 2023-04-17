package main

import (
	"errors"
	"fmt"
)

func main() {
	// NÚMEROS

	// tipos de inteiros: int8, int16, int32, int64, uint8, uint16...
	// aliases:
	// - int32 = rune
	// - uint8 = byte
	var numero int16 = 100
	fmt.Println(numero)

	// tipos de reais: float32, float64
	var numeroreal1 float32 = 1234.09
	fmt.Println(numeroreal1)

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char) // Resultado: 66 - char é convertido para o respectivo valor na tabela ascii

	// BOOLEAN
	var booleano1 bool = true
	fmt.Println(booleano1)

	// ERRO
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
