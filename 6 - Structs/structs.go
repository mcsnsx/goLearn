package main

import "fmt"

//STRUCTS é uma coleção de campos, cada campo tem um nome e um tipo, é uma estruturaa, quando não é aribuido
//um valor, o valor inicial é o valor de cada um dos campo que existem dentro da estrutura

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Maria"
	u.idade = 28
	fmt.Println(u)
	enderecoEx := endereco{"Rua dos Bobos", 0}

	u2 := usuario{"Carolina", 29, enderecoEx}
	fmt.Println(u2)

	//caso vc não tenha todos os dados do usuario
	u3 := usuario{idade: 27}
	fmt.Println(u3)

}
