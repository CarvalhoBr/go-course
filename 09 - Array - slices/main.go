package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]string
	array1[0] = "teste"

	fmt.Println(array1)

	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	// array3[6] = 3 --> array é inflexível no tamanho
	fmt.Println(array3)

	// SLICES

	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6) // slices são mais flexíveis no tamanho, porám continuam fixos no tipo
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1)) // slices são de tipos diferentes de arrays, portanto não podem ser feitas operações entre esses tipos

	slice2 := array2[1:3]
	fmt.Println(slice2)

}
