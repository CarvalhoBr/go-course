package main

import "fmt"

func sum(numbers ...int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func main() {
	fmt.Println(sum(1, 2, 4, 5, 6, 7))
}
