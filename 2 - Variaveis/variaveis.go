package main

import "fmt"

// existe várias formas de declarar variáveis, isso também se aplica as contantes, só muda "var"
// por "const"
func main() {
	var variavel1 string = "variavel 1"
	variavel2 := 2
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	//invertendo valor das variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

	const contante1 string = "constante 1"
	fmt.Println(contante1)
}
