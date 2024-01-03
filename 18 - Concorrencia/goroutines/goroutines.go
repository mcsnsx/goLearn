package main

import (
	"fmt"
	"time"
)

// CONCORRENCIA != PARALELISMO

// Paralelismo --> acontece quando você tem duas ou mais tarefas que estão sendo executadas exatamente
// ao mesmo tempo, isso só é possivel se você tiver um processador que possui mais de um núcleo, ele
// executa uma tarefa em cada núcleo, distribui as tarefas entre esses núcleos

// Concorrencia --> acontece quando as tarefas estão sendo executadas não necessáriamente ao mesmo
// tempo, elas podem ser executadas ao mesmo tempo caso você tenha um processador com mais de um
// nucleo, mas não necessáriamente. Se você tiver um processador que tem um núcleo só também é
// possivel aplicar concorrencia nele, a diferença maior é que, caso você tenha duas tarefas em
// execução, uma tarefa não esperaria a outra caabar para rodar também, de certa forma, ficariam
// revezando o tempo, rodaria um pouco da primeira tarefa e pararia, rodaria um pouco da segunda tarefa
// e parararia, ambas são executadas mas não exatamente ao mesmo tempo

func main() {
	// Goroutine --> funções ou métodos que você pode chamar a execução e não necessáriamente esperar
	// que eles terminem para continuar com o seu programa.
	go escrever("Olá Mundo!") //goroutine
	escrever("Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
