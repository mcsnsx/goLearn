package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
