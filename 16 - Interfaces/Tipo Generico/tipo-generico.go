package main

import "fmt"

// Golang é um linguagem rigida que tem grande atenção aos tipos, é preciso ter cuidado com tipos genericos
// para que isso não se torne uma gambiarra
// uma interface sem nada, seria atendida por qualquer coisa, não teria nenhuma restrição
// não é uma boa prática usar funções genericas com frequencia
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	// o "Println" é um bom exemplo de algo que recebe várias interfaces e pode funcionar com qualquer tipo
	// e qualquer quantidade/variedade de tipos, é uma função que faz sentido ter flexibilidade
	fmt.Println(1, 2, "string", false, true, float64(123456))

	//mapa sem limitação, bagunça total, isso pode ser um problema porque causa incertezas
	//recomanda-se que ela tenha algum tipo de restrição, a menos que seja um caso muito especifico
	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     "string",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
