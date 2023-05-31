package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Brandon",
		"sobrenome": "Carvalho",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Brandon",
			"ultimo":   "Carvalho",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")

	usuario2["curso"] = map[string]string{
		"nome":   "ciencia da computação",
		"campus": "uenf",
	}
	fmt.Println(usuario2)
}
