package main

import "fmt"

// closure --> é uma função como se ela tivesse uma "memoria" de onde ela veio e os valores de suas variaveis
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao

}

func main() {
	texto := "dentro da função main"
	fmt.Println(texto)

	funcNova := closure()
	funcNova()
}
