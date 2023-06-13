package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A Área da forma é %0.2f \n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return 3.14 * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 20}
	escreverArea(r)

	c := circulo{3}
	escreverArea(c)
}
