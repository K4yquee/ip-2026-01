// Escreva um programa que receba a altura de 10 atletas.
// Esse programa deve imprimir as alturas daqueles
// atletas que têm altura maior do que a média.
package main

import "fmt"

func main() {
	var vetor [10]float64
	var soma float64

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite a altura do %dº atleta: ", i+1)
		fmt.Scan(&vetor[i])

		soma += vetor[i]
	}

	media := soma / 10

	fmt.Println("Alturas acima da média:")
	for i := 0; i < 10; i++ {
		if vetor[i] > media {
			fmt.Println(vetor[i])
		}
	}
}
