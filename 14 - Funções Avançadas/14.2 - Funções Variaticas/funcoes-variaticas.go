package main

import "fmt"

// funciona como um slice, pode passar vários numeros ou nenhum
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// não é possivel ter mais de um parametro variatico por função e obrigatoriamente precisa ser o ultimo
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	//funciona como um slice, pode
	totalDaSoma := soma(1, 2, 3, 4)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 10, 2, 3, 4, 5, 6, 7)
}
