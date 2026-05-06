package main

import "fmt"

func main() {
	var vetor [10]int
	var achou int

	achou = 1

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < 10; i++ {
		if vetor[i] > 50 {
			fmt.Printf("Número %d na %dº posição da lista\n", vetor[i], i+1)
			achou = 1
		}
		if achou != 1 {
			fmt.Println("Não existe nenhum número maior que 50.")
		}
	}
}
