package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroup --> serve para sincronizar as goroutines

func main() {
	var waitGroup sync.WaitGroup

	//funciona como um grupo de espera
	waitGroup.Add(2) // é preciso especificar aqui a quantidade de goroutines que não rodar ao mesmo tempo

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() // "Done" tira 1 do contador [-1]
	}()

	go func() {
		escrever("Go")
		waitGroup.Done() // "Done" tira 1 do contador [-1]
	}()

	// Espera a contagem das goroutines chegar em zero
	waitGroup.Wait()

	// escrever("Olá Mundo!")
	// escrever("Go")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
