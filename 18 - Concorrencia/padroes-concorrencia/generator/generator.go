package main

import (
	"fmt"
	"time"
)

// Generator --> é usado para esconder a complexidade para que a função possa ser chamada normalmente na main,
// você encapsula a chamada de uma go routine e retorna um canal de comunicação que vai interagir com
// a gor routine

func main() {
	canal := escrever("Ola mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal

}
