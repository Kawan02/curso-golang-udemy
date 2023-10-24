package main

import (
	"fmt"
	"time"
)

//CONCORRÊNCIA != PARALELISMO

//PARALELISMO = Acontece quando você tem duas ou mais tarefas que estão sendo executadas exatamente ao mesmo tempo

//CONCORRÊNCIA = Elas não necessariamente estão sendo executdas ao mesmo tempo

func main() {
	go escrever("Olá Mundo!") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
