package main

import "fmt"

func main() {
	var vetor [10]int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])

		// "Empurra" o número para a posição correta
		j := i
		for j > 0 && vetor[j] < vetor[j-1] {
			vetor[j], vetor[j-1] = vetor[j-1], vetor[j] // troca
			j--
		}
	}

	fmt.Println("Vetor em ordem crescente:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", vetor[i])
	}
}
