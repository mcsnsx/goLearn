package main

import (
	"cmd/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicaçao := app.Gerar()
	if erro := aplicaçao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
