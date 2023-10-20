package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Gabriel", "Leo"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}
	// Caso não queira o indice, é só passar como _
	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, valor := range "PALAVRA" {
		fmt.Println(indice, string(valor))
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Garcia",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
