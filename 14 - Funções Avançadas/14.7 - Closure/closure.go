package main

import "fmt"

// Closure = São basicamente funções que referenciam variaveis que estão fora do seu "corpo"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()

	funcaoNova()
}
