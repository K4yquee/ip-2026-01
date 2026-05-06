package main

import "fmt"

func main() {
	var jogadas [20]int
	var freq [7]int

	for i := 0; i < 20; i++ {
		fmt.Printf("Digite o resultado da %dª jogada (1 a 6): ", i+1)
		fmt.Scan(&jogadas[i])

		if jogadas[i] >= 1 && jogadas[i] <= 6 {
			freq[jogadas[i]]++
		} else {
			fmt.Println("Valor inválido! Digite números entre 1 e 6.")
			i--
		}
	}

	fmt.Println("\nNúmeros sorteados:")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", jogadas[i])
	}

	fmt.Println("\n\nFrequência:")
	for i := 1; i <= 6; i++ {
		fmt.Printf("Número %d apareceu %d vezes\n", i, freq[i])
	}
}
