package banco

import (
	"database/sql"

	//quando tem o "_" na frente, significa um import implicito
	_ "github.com/go-sql-driver/mysql" //Driver de conexão com o MySQL
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
