package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		//time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Pedro", "Alex"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))

	}

	usuario := map[string]string{
		"nome":      "leonardo",
		"sobrenome": "silva",
	}
	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//loop infinito
	for {
		fmt.Println("executando infinitamente")
		time.Sleep(time.Second)
	}

}
