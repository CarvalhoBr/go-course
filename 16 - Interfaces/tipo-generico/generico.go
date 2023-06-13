package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("teste")
	generica(28736)
	generica(true)
}
