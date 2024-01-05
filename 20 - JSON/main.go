package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// JSON --> é um formato que pode ser usado tanto para armazenar dados como para
// transportar dados de um lugar para o outro, é muito usado quando você está mandado
// dados de um servidor para uma página web e vice versa, por exemplo. É semelhante
// ao 'struct'

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	// Aqui o JSON é printado em byte, o que é pouco prático
	fmt.Println(cachorroJson)

	// transforma um json em um struct ou map
	//json.Unmarshal()

	// ele recebe um array de bytes e converte em uma saida em JSON
	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]string{
		"nome": "Tobby",
		"raca": "Poodle",
	}

	cachorroJson2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson2)

	fmt.Println(bytes.NewBuffer(cachorroJson2))
}
