package main

import "fmt"

func main() {
	//função anonimo --> sem nome e você pode ou não armazenar o retorno dela dentro de uma variável
	retorno := func(texto string) string {
		//Sprintf --> você consegue passa uma string dentro dele concatenando com outros tipos de dados
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
		//para a funções anonima ser chamada é preciso abrir e fechar parenteses vazios
	}("Olá mundo")

	fmt.Println(retorno)

}
