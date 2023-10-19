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

}
