package auxiliar

import "fmt"

//se a primeira letra da função estiver em letra minuscula, essa função será interpretada como
//uma função privada (visisvel apenas dentro do pacote em que ela esta)
//se a primeira letra estiver maiuscula, essa função será pública (pode ser exportada)

//Boa prática --> quando uma função esta sendo exportada, ela deve ter um comentario sobre ela
//falando o que ela faz

// Escrever registra uma mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do arquivo auxiliar!")
	escrever2()
}
