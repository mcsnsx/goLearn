package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle")

	n := 10

	if n > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//quando uma variavel é criada direatmente dentro do if, ele só funciona no escopo do if, para que a variavel
	//possa ser considerada, é preciso usar o init
	if n2 := n; n2 > 0 {
		fmt.Println("Numero maior que zero")
	} else if n < -10 {
		fmt.Println("Numero que -10")
	} else {
		fmt.Println("Esta entre 0 e -10")
	}

}
