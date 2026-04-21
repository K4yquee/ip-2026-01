package main

import "fmt"

func main() {
	var preco float64
	var ingressos int
	ingressos = 130
	preco = 6.0
	despesa := 300.0

	var maxLucro, melhorPreco, lucro float64 
	var melhorIngressos int
	maxLucro = -999999

	fmt.Println("Preco\tIngressos\tLucro")

	for preco >= 1.0 {
		lucro = (preco * float64(ingressos)) - despesa

		fmt.Printf("%.2f\t%d\t\t%.2f\n", preco, ingressos, lucro)

		if lucro > maxLucro {
			maxLucro = lucro
			melhorPreco = preco
			melhorIngressos = ingressos
		}

		preco -= 0.6
		ingressos += 30
	}

	fmt.Println("\nLucro máximo:")
	fmt.Printf("Lucro: %.2f\n", maxLucro)
	fmt.Printf("Preco: %.2f\n", melhorPreco)
	fmt.Printf("Ingressos: %d\n", melhorIngressos)
}
