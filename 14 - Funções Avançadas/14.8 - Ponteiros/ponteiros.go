package main

import "fmt"

func inverterSinal(n int) int {
	return n * -1
}

// função com ponteiro
// o sinal * indica um ponteiro, ele indica onde se altera o valor de armazenamento na memoria, o * faz a
// disrreferenciação, não é preciso ter um retorno porque a alteração esta sendo feita direto na variavel
func inverterSinalPonteiro(n *int) {
	*n = *n * -1

}

func main() {
	n := 20
	numInvertido := inverterSinal(n)
	fmt.Println(numInvertido)
	fmt.Println(n)

	novoN := 40

	fmt.Println(novoN)
	//o '&' joga na função o endereço de memoria onde essa variavel esta salva
	inverterSinalPonteiro(&novoN)
	fmt.Println(novoN)
}
