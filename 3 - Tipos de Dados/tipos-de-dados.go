package main

import (
	"errors"
	"fmt"
)

func main() {

	//NUMEROS INTEIROS

	//tipos de int ---> int8, int16, int32, int64 (o numero corrsponde a quantidade de bits que
	//o numero ocupa, eles também suportam numeros negativos)

	//tipos de uint uint8, uint16, uint32, uint64 --->  = unsygned int ---> int sem sinal

	var numero int16 = 100
	fmt.Println(numero)

	//não comporta a quantidade de bits
	var numero1 int8 = 10 //100000000000
	fmt.Println(numero1)

	//se adapta de acordo com o pc
	var numero2 int = 1000000000000000
	fmt.Println(numero2)

	var numero3 uint32 = 1000 // -1000 [não funciona]
	fmt.Println(numero3)

	// alias = apelido ---> rune é exatamente igual ao int32 ; é geralmente usado quando usamos numeros que
	//representam caracteres, principalmente ta tabela ASCII
	var numero4 rune = 1234
	fmt.Println(numero4)

	//alias = apedlido ---> byte é exatamente igual a uint8
	var numero5 byte = 10
	fmt.Println(numero5)

	//NUMEROS REAIS

	//float32, float64  ---> esse numero suportam vigulas e pontos, são numeros quebrados
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//STRINGS

	//tipo explicito
	var str string = "ble ble ble"
	fmt.Println(str)

	//tipo implicito
	str2 := "texto 2"
	fmt.Println(str2)

	//ele vai printar essa informação como um numero, é o numero que o 'B' esta na tabela ASCII
	char := 'B'
	fmt.Println(char)

	//todo tipo tem seu valor inicial, quando algo esta sendo declarado mas não é atribuido um valor

	//variavel valor inicial -- vazio
	var texto string
	fmt.Println(texto)

	//variavel valor inicial -- zero
	var texto1 int
	fmt.Println(texto1)

	//tipo booleano [bool] -- verdadeiro ou falso, true or false, valor inicial de bool é false
	var booleano1 bool = true
	fmt.Println(booleano1)

	//tipo error --> erro em go é tratado como um tipo, valor inicial <nil>
	// <nil> --> serve como valor inicial de muitos tipos de dados
	// para criar um erro em go é preciso usar um pacote "errors"
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
