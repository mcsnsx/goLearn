package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	// caso vc queira ignorar uma das informações para que ela não seja convertida,
	// basta colocar um tracinho no campo
	// ex.: Nome  string `json:"-"`
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// caso vc queira ignorar uma das informações, basta colocar um tracinho no campo
	cachorroJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`

	var c cachorro

	// é preciso passar os dados pra dentro do struct cachorro, usamos unmarsha
	// usamos dois parametros: os dados e o endereço de memoria onde vamos jogar
	// esses dados, usamos o ponteiro porque queremos que essa variavel fique
	// alterada depois dessa linha
	// o parametro recebido por ele deve ser um slice de byte,
	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	// convertendo pra map, é preciso se atentar ao tipo de dado
	cachorroJSON2 := `{"nome":"Tobby","raca":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroJSON2), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
