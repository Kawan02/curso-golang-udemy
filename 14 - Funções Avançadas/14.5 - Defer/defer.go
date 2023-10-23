package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcão 1")
}

func funcao2() {
	fmt.Println("Executando a funcão 2")

}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função, para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media < 6 {
		return false
	}
	return true
}

func main() {
	// DEFER = ADIAR
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
