package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// Interface --> estrutura abstrata
// A imp´lementação de interfaces é implicita, não precisa ser declarada na mão assim como em outras
// linguagens orientadas a objeto
// como go é um linguagem rigida com o tipos, a interface deixa isso mais flexivel
type forma interface {
	area() float64
}

// usando interface você tem uma função que ao invés de receber uma estrutura concreta, ela recebe uma
// estrutura abstrata que é o caso da interface
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

// Para satisfazer essa ihterface, é preciso implementar um método área
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	//return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}

	//não é possivel chamar a função e a forma diretamente porque dessa forma ele não atende os requisitos
	//para ser uma forma (precisa necessáriamente ter um método, nesse caso, chaamado área que não recebe
	// parametros e retorne um float64), se eu quiser passar qualquer coisa para a função, precisa ter um
	// método onde a assinatura é igual a que esta no metodo, o método faz o que está dentro das chaves dele
	// e pode variar, mas a assinatura precisa ser igual
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
