package main

import "fmt"

//funções recursivas ---> são funções que chamam a elas mesmas, para ela funcionar, ela precisa de uma outra execução
//dela mesma. Essas funções tem uma condição de parada para não cair em um looping infinito e não acontecer o
//"estouro de pilha", as funções recursivas não devem ser usadas em situações com muitas execussões

// sequencia fibonacci --> é uma sdquencia de numeros onde o proximo numero é sempre a soma dos dois
// numeros anteriores (1 1 2 3 5 8 13 ...)
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(12)

	for i := uint(1); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	//fmt.Println(fibonacci(posicao))

}
