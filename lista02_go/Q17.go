package main

import "fmt"

func main() {
	var conta int
	var tipo string
	var consumo, valor float64

	fmt.Print("Digite o número da conta: ")
	fmt.Scan(&conta)

	fmt.Print("Digite o tipo (R - residencial, C - comercial, I - industrial): ")
	fmt.Scan(&tipo)

	fmt.Print("Digite o consumo de água (m3): ")
	fmt.Scan(&consumo)

	if tipo == "R" || tipo == "r" {
		valor = 5 + (consumo * 0.05)

	} else if tipo == "C" || tipo == "c" {
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + (consumo-80)*0.25
		}

	} else if tipo == "I" || tipo == "i" {
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + (consumo-100)*0.04
		}

	} else {
		fmt.Println("Tipo inválido!")
		return
	}

	fmt.Println("Conta:", conta)
	fmt.Printf("Valor a pagar: R$ %.2f\n", valor)
}
