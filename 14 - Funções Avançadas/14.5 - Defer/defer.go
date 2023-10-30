package main

import "fmt"

func func1() {
	fmt.Println("Executando a função 1")
}

func func2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	fmt.Println("Função aluno aprovado")
	defer fmt.Println("Medica calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

// Defer --> é uma claudula utilizada quando precisamos adiar a execução de um determinado pedaço de código, pode ser
// uma função criada ou de algum pacote,
// Defer = Adiar
func main() {
	//defer func1()
	//func2()
	fmt.Println(alunoAprovado(7, 8))

}
