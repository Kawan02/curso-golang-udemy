package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	// Enviando um valor para o CANAL
	canal <- "Olá Mundo!"
	canal <- "Programando em GO!"

	// A variavel mensagem está recebendo o valor que está no CANAL
	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
