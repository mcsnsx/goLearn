package main

import "fmt"

func main() {
	//ARITMETICOS ---> + - * / %
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//as operações devem ser feitas com dados dos mesmo tipo
	var numero1 int16 = 10
	var numero2 int16 = 10
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)

	//ATRIBUIÇÃO ---> = :=
	var variavel1 string = "string 1"
	variavel2 := "string 2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS ---> > < <= >= != == === e etc, dempre retornam um bool
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)
	fmt.Println("----------------------------------------------------")

	//OPERADORES LÓGICOS ---> &&, || e  !
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNARIOS ++, --, +=, -=
	numero3 := 10
	numero3++
	numero3 += 15
	numero3--
	numero3 -= 20
	numero3 *= 3
	numero3 /= 10
	numero3 %= 3
	fmt.Println(numero3)

	//OPERAODR TERNARIO ---> ele tem uma condições a ser vaaliada, um valor para a variavel caso ela seja
	//verdadeira e um valor para a variavel caso ela seja falsa.

	//texto := numero3 > 5 ? "Maior que 5" : "Menor que 5"
	//fmt.Println(texto)

	// ISSO NÃO EXISTE EM GO, é preciso fazer um if e else

	var texto string
	if numero3 > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
