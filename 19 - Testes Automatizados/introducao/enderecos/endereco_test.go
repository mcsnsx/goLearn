package enderecos

import "testing"

//Testing --> o teste nunca fica no mesmo arquivo da função que vai ser testada
// porque os arquivos de teste precisam ter um nome especifico para que o go entenda
// ('nomedoarquivo_test.go'), e nesse caso, isso vai ser rodado por linha de comando
// e esse comando encontra os arquivos que tem '_test.go' e executa as funções de teste
// que existem dentro deles

//Teste de Unidade --> testa a menor unidade de um código, no caso, vai ser uma função.
// Os testes são a unica excessão onde podemos ter mais de um pacote dentro da mesma
// pasta/diretorio

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// O nome da função de teste tem que começar com 'Test' e o nome da função que será testada
// ex: TestXxxxxxx, esse função deve receber um ponteiro como parametro, do tipo T
func TestTipoDeEndereco(t *testing.T) {
	// enderecoParaTeste := "Rua ABC"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoEndereco(enderecoParaTeste)

	// faz os testes rodarem em paralelo
	t.Parallel()

	// essa é uma forma de testar várias respostas possiveis em um mesmo slice
	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		//{"Praça das Rosas", "tipo invalido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		//{"", "tipo invalido"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}

	// Se ele cair aqui, é dito que o "teste quebrou", se ele não cair aqui, é porque
	// o "teste passou"
	// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	// 	}
}

// Esse teste não é considerado muito confiável porque está testando apenas um caso,
// apenas uma variável, mas é possivel fazer inumeros testes com o mesmo teste

func TestQualquer(t *testing.T) {
	// faz os testes rodarem em paralelo
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
