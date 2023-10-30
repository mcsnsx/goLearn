package main

import "fmt"

var n int

// vai ser executada antes da função main, mesmo que esteja em outra ordem, podemos ter uma função init por arquivo
// e não uma por pacote igual a main, e a função init é sempre executada primeiro, isso serve para inicializar
// algum setup ou alguma variavel que sera usada pela função ou de forma global
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função Main sendo executada")
	fmt.Println(n)

}
