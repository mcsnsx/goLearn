package main

import "fmt"

// a diferença entre metodo e função é que o metodo ta sempre associado a alguma coisa (struct, interface, etc)
// cria-se metodos para dar ações as funções, é algo parecido com a programação orientada a objetos onde vc
// tem as classes, os atributos e os objetos
type usuario struct {
	nome  string
	idade uint8
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)

}
