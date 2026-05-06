// Escreva um programa que receba quinze números inteiros e armazene em um vetor
// a raiz quadrada de cada
// número. Caso o valor digitado seja menor do que zero, o número -1 deve ser
// atribuído ao elemento do vetor.
// Após isso, imprima todos os valores armazenados.
package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	var vetor [15]float64

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&n)

		if n < 0 {
			vetor[i] = -1
		} else {
			vetor[i] = math.Sqrt(n)
		}
	}

	fmt.Println("Valores armazenados:")
	for i := 0; i < 15; i++ {
		fmt.Printf("%.1f\n", vetor[i])
	}
}
