package main

// Worker Pools --> é como se você tivesse uma fila de tarefaas para serem executadas e você tem vários workers
// vários processos pegando itens dessa fila de maneira independente

func main() {

}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
