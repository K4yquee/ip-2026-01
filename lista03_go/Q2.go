package main

import "fmt"

func main() {
	var soma, quantidade int
	var media float64

	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			soma = soma + i
			quantidade++
		}
	}

	media = float64(soma) / float64(quantidade)

	fmt.Printf("A soma dos valores e = %d\n", soma)
	fmt.Printf("A média da soma dos valores e %.2f", media)
}
