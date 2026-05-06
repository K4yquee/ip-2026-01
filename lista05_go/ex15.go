//Crie um programa que leia um vetor de 30 números inteiros e gere um segundo
//vetor cujas posições pares
//possuirão elementos que serão o dobro do vetor original e as ímpares, o triplo.

package main

import "fmt"

func main() {
	var vetor1 [30]int
	var vetor2 [30]int

	for i := 0; i < 30; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			vetor2[i] = vetor1[i] * 2
		} else {
			vetor2[i] = vetor1[i] * 3
		}
	}

	fmt.Println("Vetor resultante:")
	for i := 0; i < 30; i++ {
		fmt.Printf("%dº posição : %d\n", i, vetor2[i])
	}
}
