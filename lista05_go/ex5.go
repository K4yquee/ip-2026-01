package main

import "fmt"

func main() {
	var vetor [10]int
	var a int
	var b int

	b = 0

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])

	}

	a = vetor[0]
	b = 0

	for i := 1; i < 10; i++ {
		if vetor[i] < a {
			a = vetor[i]
			b = i
		}
	}

	fmt.Printf("O menor número é %d, na %dº posição \n", a, b+1)
}
