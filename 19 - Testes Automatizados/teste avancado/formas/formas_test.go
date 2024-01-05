package formas

import (
	"math"
	"testing"
)

// TDD - Test Driven Developer --> quando você vai desenvolver uma funcionalidade, você faz o
// teste primeiro, principalmente quando usa casos como esse com o uso de interfaces e sugere
// calculos, há uma série de particularidades e coisas que você aplica para essa pratica
// (fazer o teste primeiro imaginando o que a função vai fazer escrevendo o minimo de codigo
// possivel para a função passar o seu teste e somente depois vai incrementando a função
// de acordo com a evolução do teste)

func TestArea(t *testing.T) {
	// usa-se 't.Run' quando você tem duas funções que vão representar dois netodos que estão
	// sendo testados, é uma forma mais organizada de separar os testes
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			// Fatalf --> funciona como o 'Errorf' só que ele vai parar aqui, não vai proceguir para
			// as demais
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
