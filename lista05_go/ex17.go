//Escreva um programa que leia um vetor de dez elementos inteiros e mostre os números primos e suas
//respectivas posições.

package main

import "fmt"

func main() {
	var vetor [10]int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])
	}

	fmt.Println("Números primos e suas posições:")

	encontrou := false

	for i := 0; i < 10; i++ {
		n := vetor[i]

		if n <= 1 {
			continue
		}

		primo := true

		for j := 2; j < n; j++ {
			if n%j == 0 {
				primo = false
				break
			}
		}

		if primo {
			fmt.Printf("Número %d na posição %d\n", n, i)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nenhum número primo encontrado.")
	}
}
