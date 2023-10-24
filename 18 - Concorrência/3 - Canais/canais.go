package main

import (
	"fmt"
	"time"
)

func main() {
	//CANAL
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	//Opção 1: Validamos se o canal está aberto, caso não esteja, ele sairá do for
	for {
		//(<-CANAL) Recebendo o valor. Estou esperando que o CANAL receba um valor
		mensagem, aberto := <-canal
		if !aberto {
			//Break = É uma maneira de sair do for infinito
			break
		}
		fmt.Println(mensagem)
	}

	//Opção 2: É uma maneira mais usada e recomendada para resolver o erro quando o canal estiver fechado
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//(CANAL <-) Passando o valor para o CANAL
		canal <- texto
		time.Sleep(time.Second)
	}

	// Função nativa do GO, que fechará o CANAL após a execução do for
	close(canal)
}
