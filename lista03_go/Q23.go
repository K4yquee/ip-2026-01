package main

import "fmt"

func main() {
	var n int
	var soma float64 = 0

	fmt.Print("Digite o valor de N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Valor inválido!")
		return
	}

	for i := 1; i <= n; i++ {
		numerador := 1000 - (i-1)*3
		termo := float64(numerador) / float64(i)

		if i%2 == 0 {
			soma -= termo
		} else {
			soma += termo
		}
	}

	fmt.Printf("Resultado da série: %.2f\n", soma)
}
