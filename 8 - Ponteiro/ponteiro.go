package main

import "fmt"

func main() {
	//passar valores por copia
	var var1 int = 10
	var var2 = var1
	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	//Ponteiro - referencia de memoria (endereço de memoria onde a informação fica salva)
	var var3 int
	var ponteiro *int
	var3 = 100
	ponteiro = &var3

	fmt.Println(var3, *ponteiro) //desreferenciação -- desfazer a referencia

	var3 = 150
	fmt.Println(var3, ponteiro)
}
