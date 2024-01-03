package main

import (
	"fmt"
	"time"
)

// Select --> é muito parecido com o switch, faz com que não haja codependencia de execução entre canais
// isso ajuda a ganhar tempo e até perder um pouco da concorrencia visto que os dados podem ter tempos
// de execução diferentes entre si, geralmente é usado em concorrencias e funções com tempos diferentes

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

		// mensagemCanal1 := <-canal1
		// fmt.Println(mensagemCanal1)

		// mensagemCanal2 := <-canal2
		// fmt.Println(mensagemCanal2)
	}

}
