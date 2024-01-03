package main

import "fmt"

// operações de receber e enviar dados são operações bloquiantes, quando usado na mesma função podem
// bloquear algum dado / fluxo então se algo bloquear esse dado e ele não for recebido onde ele deve
// ser recebido causa o deadlock, é por esse motivo que os canais geralmente são usados em funções
// separadas, uma função que envia o valor e outra que recebe o valor. A alterantiva para poder usar
// dentro da mesma função é criar um canal com buffer onde você especifica a capacidade para o seu
// canal

// Canal com buffer --> ele só vai bloquear quando atingir a capacidade máxima dele

func main() {
	canal := make(chan string, 2) // o 2 é a capacidade do canal
	canal <- "Ola mundo"

	mensagem := <-canal
	fmt.Println(mensagem)

}
