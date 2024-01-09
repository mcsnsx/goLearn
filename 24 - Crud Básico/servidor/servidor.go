package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// nós precisamos configurar a função pra ler o corpo da requisição para que ele entenda
// a entrada e saida dos dados, a função em si vai inserir um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	// Como seria a query no banco:
	// INSERT INTO USUARIOS (nome, email) values ("nome", "email")
	// Nesse caso usaremos o 'PREPARE STATEMENT' que é muito usado para evitar um ataque
	// que chama 'sql injection' que é uma ação maliciosa para executar mais de um comando
	// no seu banco de dados
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	//Status Codes --> é o que indica o que aconteceu com a requisição
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}
