package main

import "fmt"

func diaDaSemana(n int) string {
	switch n {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "numero invalido"
	}
}

// outra forma de sintax do switch
func diaDaSemana2(n int) string {
	var diaDaSemana string
	switch {
	case n == 1:
		diaDaSemana = "Domingo"
		fallthrough // faz ele apontar par a o proximo case
	case n == 2:
		diaDaSemana = "Segunda"
	case n == 3:
		diaDaSemana = "Terça"
	case n == 4:
		diaDaSemana = "Quarta"
	case n == 5:
		diaDaSemana = "Quinta"
	case n == 6:
		diaDaSemana = "Sexta"
	case n == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "numero invalido"
	}

	return diaDaSemana

}

func main() {
	fmt.Print("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
