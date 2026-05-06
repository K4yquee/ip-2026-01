package main

import "fmt"

func main() {
	var vetor [10]int
	var pares [10]int
	var impares [10]int

	var somapar int
	var qimpar int

	j := 0
	k := 0

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d número: ", i+1)
		fmt.Scan(&vetor[i])

		if vetor[i]%2 == 0 {
			somapar += vetor[i]
			pares[j] = vetor[i]
			j++
		} else {
			impares[k] = vetor[i]
			k++
			qimpar++
		}
	}

	fmt.Print("Números pares: ")
	for i := 0; i < j; i++ {
		fmt.Print(pares[i], " ")
	}

	fmt.Println("\nSoma dos pares:", somapar)

	fmt.Print("Números ímpares: ")
	for i := 0; i < k; i++ {
		fmt.Print(impares[i], " ")
	}

	fmt.Println("\nQuantidade de ímpares:", qimpar)
}
