package main

import "fmt"

func main() {
	var x, soma, fatorial float64

	fatorial = 1

	fmt.Print("Digite o valor de X: ")
	fmt.Scan(&x)

	soma = x

	for i := 1; i <= 19; i++ {
		fatorial *= float64(i)

		termo := x / fatorial

		if i%2 != 0 {
			soma -= termo
		} else {
			soma += termo
		}
	}

	fmt.Printf("Resultado da soma: %.4f\n", soma)
}
