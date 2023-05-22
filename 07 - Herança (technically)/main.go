package main

import "fmt"

type person struct {
	nome   string
	idade  uint8
	altura uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	p1 := person{"João", 20, 160}
	fmt.Println(p1)

	s1 := student{p1, "Computação", "UENF"}
	fmt.Println(s1)
}
