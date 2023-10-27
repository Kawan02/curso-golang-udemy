package main

import (
	"log"
	"net/http"
)

// HTTP => É um protocolo de comunicação - Base da comunicação WEB
// Cliente - Servidor

func paginaRaiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {

	http.HandleFunc("/", paginaRaiz)
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
