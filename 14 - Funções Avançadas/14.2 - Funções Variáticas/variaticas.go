package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, valor := range numeros {
		total += valor
	}
	return total
	// fmt.Println(numeros)
}

func main() {
	totalDaSoma := soma(10, 20, 30, 40, 50)

	fmt.Println(totalDaSoma)
}
