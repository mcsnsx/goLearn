package main

import "fmt"

// Worker Pools --> é como se você tivesse uma fila de tarefaas para serem executadas e você tem vários workers
// vários processos pegando itens dessa fila de maneira independente

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	//para diminuir o tempo de execução, podemos chamar as goroutines mais de uma vez
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// nesse caso é possivel especificar se o canal envia (chan<-) ou recebe (<-chan) dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
