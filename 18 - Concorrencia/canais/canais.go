package main

import (
	"fmt"
	"time"
)

// Canais --> tem esse nome porque é um canal de comunicação, pode enviar e receber dados e pode ser
// usado para sincronizar as goroutines, os canais tem duas operações que são basicas e muito importantes
// que é enviar e receber um dado, essas operações bloqueiam a execução do programa
// Os canais tem a propriedade de estar aberto ou fechado, se ele esta aberto significa que ele ainda vai
// enviar dadaos então pode receber dados também, se eles esta fechado então ele não envia e nem recebe dados

// deadlock --> é quando você não tem mais nenhum lugar que esta enviando dado para o seu canal só
// que o seu canal ainda ta esperando receber um dado, seu programa fica eternamente esperando um dado que
// nunca vai chegar e por isso ele trava, geralmente não é pego em compilação, o go não consegue identificar
// ele só pega em execução

func main() {
	// o tipo estabelecido após o 'chan' significa que só poderam ser trafegados dados desse tipo
	canal := make(chan string)

	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// o canal recebe um valor e salva o valor na variavel mensagem, é nesse ponto que a sincronização
	// com a goroutine acontece
	for {
		//é possivel passar a mensagem e um retorno de verificação para saber se o canal esta aberto
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	//não é preciso fazer um looping infinito pra essa ideia dos canais funcionar, existe essa segunda maneira:
	// esse trecho cria a condição de que para cada mensagem recebida no canal enquanto ele estiver aberto
	// ele vai printar na tela, o que facilita a sintaxe
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// envia um dado por um canal
		canal <- texto
		time.Sleep(time.Second)
	}
	//fecha o canal para ele não possa mais enviar e nem receber dados
	close(canal)
}
