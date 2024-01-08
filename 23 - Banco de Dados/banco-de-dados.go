package main

import (
	"database/sql"
	"fmt"
	"log"

	//quando tem o "_" na frente, significa um import implicito
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// devemos passar uma string de conexão com os dados de usuario e senha de acesso
	// formato padrão de string de conexão "usuario:senha@/nomedobanco"
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	// abrimos a conexã com o banco
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("Dentro do SQL.Open")
		log.Fatal(erro)
	}

	// aqui fechamos a conexão com o banco
	defer db.Close()

	// aqui damos um ping no banco para testar a conexão
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta")

	// Podemos fazer uma query no banco direto do golang
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
