package main

import "fmt"

func main() {
	// Função anonima
	func() {
		fmt.Println("Olá mundo")
	}()

	// Função anonima com parâmetro
	func(texto string) {
		fmt.Println(texto)
	}("Passando parâmetro")

	// Função anonima com retorno
	retorno := func(texto string) string {
		// No Print representa uma string (%s)
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
