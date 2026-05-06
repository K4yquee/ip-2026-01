package main

import "fmt"

func main() {
	var vetor [10]int
	var vetor2 [5]int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d número do primeiro vetor: ", i+1)
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Digite o %d número do segundo vetor: ", i+1)
		fmt.Scan(&vetor2[i])
	}

	somaVetor2 := 0
	for i := 0; i < 5; i++ {
		somaVetor2 += vetor2[i]
	}

	var rPar [10]int
	var rImpar [10]int

	j := 0
	k := 0

	for i := 0; i < 10; i++ {
		if vetor[i]%2 == 0 {
			rPar[j] = vetor[i] + somaVetor2
			j++
		} else {
			rImpar[k] = vetor[i] + somaVetor2
			k++
		}
	}

	fmt.Println("\nVetor resultado (pares):")
	for i := 0; i < j; i++ {
		fmt.Print(rPar[i], " ")
	}

	fmt.Println("\nVetor resultado (ímpares):")
	for i := 0; i < k; i++ {
		fmt.Print(rImpar[i], " ")
	}
}
