package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func message(w http.ResponseWriter, message string) {
	w.Write([]byte(message))
}

func messageServer(w http.ResponseWriter, i int) {
	w.WriteHeader(i)
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Falha ao ler o corpo da requisição!")
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Erro ao converter o usuário para struct")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messageServer(w, http.StatusInternalServerError)
		message(w, "Erro ao conectar no banco de dados!")
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		messageServer(w, http.StatusBadRequest)
		// w.Write([]byte(("Erro ao criar o statement!")))
		message(w, "Erro ao criar o statement!")
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		messageServer(w, http.StatusBadRequest)
		// w.Write([]byte(("Erro ao executar o statement!")))
		message(w, "Erro ao executar o statement!")
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		messageServer(w, http.StatusNotFound)
		message(w, "Erro ao obter o ID inserido!")
		// w.Write([]byte(("Erro ao obter o ID inserido!")))
		return
	}

	// w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
	messageServer(w, http.StatusCreated)
	message(w, fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido))
}

// BuscarUsuarios traz todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		messageServer(w, http.StatusInternalServerError)
		message(w, "Falha ao conectar com o banco de dados!")
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		message(w, "Erro ao buscar os usuários")
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			message(w, "Erro ao escanear o usuário")
			return
		}
		usuarios = append(usuarios, usuario)
	}

	// Transforma o slice de usuarios em JSON
	messageServer(w, http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		message(w, "Erro ao converter os usuários para JSON")
		return
	}
}

// BuscarUsuario traz um usuários específico salvo no banco de dado
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)
	if erro != nil {
		message(w, "Erro ao converter o parâmetro para inteiro")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messageServer(w, http.StatusInternalServerError)
		message(w, "Erro ao conectar com o banco de dados!")
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		messageServer(w, http.StatusNotFound)
		message(w, "Erro ao buscar o usuário!")
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			messageServer(w, http.StatusNotFound)
			message(w, "Erro ao escanear o usuário!")
			return
		}
	}

	if usuario.ID == 0 {
		messageServer(w, http.StatusNotFound)
		message(w, "ID não encontrado!")
		return
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Erro ao converter o usuário para JSON!")
		return
	}
}

// AtualizarUsuario altera os dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)
	if erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Erro ao converter o parâmetro para inteiro!")
		return
	}

	corpodaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Erro ao ler o corpo da requisição!")
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpodaRequisicao, &usuario); erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Erro ao converter um usuário para struct")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messageServer(w, http.StatusInternalServerError)
		message(w, "Erro ao conectar no banco de dados!")
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Erro ao criar o statement!")
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		message(w, "Erro ao atualizar o usuário!")
		return
	}
	messageServer(w, http.StatusNoContent)
}

// DeleterUsuario exclui um usuário específico no banco de dados
func DeleterUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)
	if erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Erro ao converter o parâmetro para inteiro!")
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messageServer(w, http.StatusInternalServerError)
		message(w, "Erro ao conectar no banco de dados!")
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		messageServer(w, http.StatusBadRequest)
		message(w, "Erro ao criar o statement!")
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		message(w, "Erro ao deletar o usuário!")
		return
	}

	messageServer(w, http.StatusNoContent)
}
