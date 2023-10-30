package main

import "fmt"

func recuperarExecucao() {
	//fmt.Println("Tentativa de recupera a execução")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}

}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	fmt.Println("Função aluno aprovado")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	//a função panic interrompe o fluxo normal do seu programa quando retorna algo inesperado e vai parar o processo, ele vai "entrar em panico"
	//ela vai chamar todas as funções que tem defer se você não chamar o recover e o programa morre
	//o erro pode ser tratado e continuar a execussão do programa, o panic interrompe e mata o programa e nada
	//que vem a seguir é executado, o programa só pode ser recuperado usando a clausula recover, não é recomendado
	//para tratamento de erros
	panic("A media é exatamente 6!")
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução")
}
