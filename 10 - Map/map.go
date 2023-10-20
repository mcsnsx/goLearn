package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	//estrutura semelhante ao struckt, porem, todas as chaves tem que ter o mesmo tipo de dados e esse tipo
	//deve ser declarado, deve manter os tipos coerentes
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	//para chamar uma caracteristica especifica é preciso chamar o nome da chave
	fmt.Println(usuario["nome"])

	//usuario com map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "campus 1",
		},
	}

	fmt.Println(usuario2)

	//adicionando uma chave
	usuario2["signo"] = map[string]string{
		"nome": "gemeos",
	}
	fmt.Println(usuario2)

	//deletando uma chave
	delete(usuario2, "campus")
	fmt.Println(usuario2)

}
