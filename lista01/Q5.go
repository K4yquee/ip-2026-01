package main

import "fmt"

func main() {
	var conta int
	var consumo, valor float64
	var tipo string

	fmt.Println("Digite o numero da conta: ")
	fmt.Scan(&conta)

	fmt.Println("Digite a quantia consumida de agua: ")
	fmt.Scan(&consumo)

	fmt.Println("Digite o tipo de consumidor: ")
	fmt.Scan(&tipo)

	if tipo == "R" {
		valor = 5 + (0.05 * consumo)
	} else if tipo == "C" {
		valor = 500 + (0.25 * consumo)
	} else if tipo == "I" {
		valor = 800 + (0.04 * consumo)
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}
