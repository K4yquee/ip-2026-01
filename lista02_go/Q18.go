package main

import "fmt"

func main() {
	var preco, valor float64
	var dia int
	var categoria string

	fmt.Print("Digite o preço normal do DVD: ")
	fmt.Scan(&preco)

	fmt.Print("Digite o dia da semana (1 a 7): ")
	fmt.Scan(&dia)

	fmt.Print("Digite a categoria (C - comum, L - lançamento): ")
	fmt.Scan(&categoria)

	valor = preco

	if dia == 2 || dia == 3 || dia == 5 {
		valor = valor * 0.6
	}

	if categoria == "L" || categoria == "l" {
		valor = valor * 1.15
	}

	fmt.Printf("Valor final: R$ %.2f\n", valor)
}
