package main

import "fmt"

// MÉTODOS

// a diferença entre metodo e função é que o metodo ta sempre associado a alguma coisa (struct, interface, etc)
// cria-se metodos para dar ações as funções, é algo parecido com a programação orientada a objetos onde vc
// tem as classes, os atributos e os objetos
// diferente da função, o método precisa estar atrelado a alguma estrutura, ele faz parte da estrutura, não é algo
// que pode ficar solto
type usuario struct {
	nome  string
	idade uint8
}

// cria-se uma função que esta vinculada a uma certa estrutura (nesse caso struct usuario), isso significa
// que todos os usuarios tem a função de salvar
// o "u" é a váriavel que usamos para referenciar outros campos do mesmo usuario que chamou esse método
func (u usuario) salvar() {
	fmt.Println("Dentro do método salvar")
	// string 		-> %s
	// numero 		-> %d
	// numero real  -> %f
	// pula linha 	-> \n
	// após o termino da string, você passa a string com o valor que vai substituir o '%s'
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

// os métodos seguem a mesma lógica das funções, pode ter retorno, parametros e tudo mais
func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

// os métodos podem alterar valores
// o '*' é um ponteiro
// quando você tem um método que vai fazer uma alteração em algum campo do struct, você passa a referencia
// do struct como um ponteiro, quando você não vai alterar o valor, não há necessidade de passar com um ponteiro
func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	//fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorIdade()
	fmt.Println(maiorDeIdade)

	usuario2.aniversario()
	fmt.Println(usuario2.idade)
}
