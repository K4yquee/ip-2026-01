package main

import "fmt"

func main() {
	var precoInicial, precoFinal float64
	var ar, pintura, vidro, direcao string

	fmt.Print("Digite o preço inicial do carro: ")
	fmt.Scan(&precoInicial)

	precoFinal = precoInicial

	fmt.Print("Deseja ar condicionado? (s/n): ")
	fmt.Scan(&ar)
	if ar == "s" {
		precoFinal += 1750
	}

	fmt.Print("Deseja pintura metálica? (s/n): ")
	fmt.Scan(&pintura)
	if pintura == "s" {
		precoFinal += 800
	}

	fmt.Print("Deseja vidro elétrico? (s/n): ")
	fmt.Scan(&vidro)
	if vidro == "s" {
		precoFinal += 1200
	}

	fmt.Print("Deseja direção hidráulica? (s/n): ")
	fmt.Scan(&direcao)
	if direcao == "s" {
		precoFinal += 2000
	}

	fmt.Println("Preço final do carro: R$", precoFinal)
}
