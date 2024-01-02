package app

import "github.com/urfave/cli"

// Essa funçaõ retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	return app
}
