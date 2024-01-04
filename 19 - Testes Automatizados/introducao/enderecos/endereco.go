package enderecos

import "strings"

// TipoEndereco verifica se o endereço tem um tipo válido e o retorna
func TipoEndereco(endereco string) string {
	//declaração dos tipos válidos
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	//conversão para letras minusculas
	enderecoMinusculo := strings.ToLower(endereco)

	// a função split divide a string em um slice, usa o espaço em branco para distinguir uma palavra
	// da outra e ve se a primeira palavra bate com algum dos tipos
	primeiraPalavra := strings.Split(enderecoMinusculo, " ")[0]

	enderecoValido := false
	//se algum dos itens do slice bater, ele vira true
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoValido = true
		}
	}
	if enderecoValido {
		return strings.Title(primeiraPalavra)
	}
	return "tipo invalido"
}
