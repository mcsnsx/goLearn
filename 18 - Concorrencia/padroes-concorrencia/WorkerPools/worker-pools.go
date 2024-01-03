package main

// Worker Pools --> é como se você tivesse uma fila de tarefaas para serem executadas e você tem vários workers
// vários processos pegando itens dessa fila de maneira independente

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
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
