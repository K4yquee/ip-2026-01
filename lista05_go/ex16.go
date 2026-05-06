//Dado um vetor com dados de 50 idades, elabore um programa que permita calcular a moda das idades. Obs.:
//Moda é o valor que tem maior incidência de repetições.

package main

import "fmt"

func main() {
	var vetor [50]int

	for i := 0; i < 50; i++ {
		fmt.Printf("Digite a idade da %dº pessoa: ", i+1)
		fmt.Scan(&vetor[i])
	}

	moda := vetor[0]
	mc := 0

	for i := 0; i < 50; i++ {
		c := 0

		for j := 0; j < 50; j++ {
			if vetor[i] == vetor[j] {
				c++
			}
		}

		if c > mc {
			mc = c
			moda = vetor[i]
		}
	}

	fmt.Printf("A moda das idades é: %d\n", moda)
	fmt.Printf("Ela aparece %d vezes\n", mc)
}
