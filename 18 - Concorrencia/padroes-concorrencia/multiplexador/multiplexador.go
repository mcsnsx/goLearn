package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Multiplexador --> é quando vc pega dois ou mais canais e junta em um canal só, isso centraliza a
// comunicação em um luhar só

func main() {
	canal := multiplexar(escrever("ola"), escrever("bye"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Essa função recebe dois canais como entrada e retorna um canal só como saida
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			// esse select vai ver qual dos dois canais tem uma mensagem disponivel pra ser lida, ele
			// faz uma especie de encaminhamento de mensagem
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			// deixando o tempo pseudo aleatorio
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal

}
