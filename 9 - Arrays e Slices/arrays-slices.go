package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//Array --> listas manpulaveis
	//array sem todas as posições preenchidas (retorna espaços em branco)
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	//array com todas as posições preenchidas
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	//array sem uma quantidade de posições definidas, o tamanho yda lista é determinado a partir da
	//quantidade de valores passados
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//Slice (fatia) --> estrutura baseada no array, porem com algumas flexibilidades a mais, principalmente o tamanho
	//não tem um tamanho fixo, pode mudar de acordo com a necessidade, tem tamanho dinamico..
	//ele não é um array, ele aponta para um array
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	//append --> adiciona um item no slice e retorna um slice novo com esse item adicionado
	slice = append(slice, 18)
	fmt.Println(slice)

	//referenciar array --> o slice sera um pedaço do array
	slice2 := array2[1:3] //ponteiro --> aponta para as posições desejadas
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	fmt.Println("-------------------------------------------------------------------------")

	//ARRAYS INTERNOS -->
	//função make --> aloca um espaço na memoria para uma determinada coisa
	//sempre que a capacidade de um slice é estourada, ele dobra o seu tamanho pra que essa capacidade nunca
	//seja limitada
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	fmt.Println(len(slice3)) // len (length) --> função para ver o tamanho
	fmt.Println(cap(slice3)) // cap --> função para ver a capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	//quando não deficimos a capacidade do slice na função make, ele é definido como o tamanho dele
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
