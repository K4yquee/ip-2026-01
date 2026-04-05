package main

import "fmt"

func main() {
	var dia, mes, ano int
	var nomeMes string

	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)

	fmt.Print("Digite o mês (1-12): ")
	fmt.Scan(&mes)

	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)

	switch mes {
	case 1:
		nomeMes = "janeiro"
	case 2:
		nomeMes = "fevereiro"
	case 3:
		nomeMes = "março"
	case 4:
		nomeMes = "abril"
	case 5:
		nomeMes = "maio"
	case 6:
		nomeMes = "junho"
	case 7:
		nomeMes = "julho"
	case 8:
		nomeMes = "agosto"
	case 9:
		nomeMes = "setembro"
	case 10:
		nomeMes = "outubro"
	case 11:
		nomeMes = "novembro"
	case 12:
		nomeMes = "dezembro"
	default:
		fmt.Println("Mês inválido!")
		return
	}

	fmt.Printf("%d de %s de %d\n", dia, nomeMes, ano)
}
