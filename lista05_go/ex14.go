//Faça um programa que leia dois vetores de 10 elementos inteiros cada um e mostre
//o vetor resultante da
//intercalação desses dois vetores.
//Exemplo: Vetor 1 [0 5 4 2 1 5 3 2 5 9]
//Vetor 2 [1 5 4 2 0 5 3 2 5 9]
//Vetor resultante da intercalação [0 1 5 5 4 4 2 2 1 0 5 5 3 3 2 2 5 5 9 9]

package main

import "fmt"

func main() {
	var v1, v2 [10]int
	var v3 [20]int

	// leitura do vetor 1
	for i := 0; i < 10; i++ {
		fmt.Printf("Vetor 1 - Digite o %dº valor: ", i+1)
		fmt.Scan(&v1[i])
	}

	// leitura do vetor 2
	for i := 0; i < 10; i++ {
		fmt.Printf("Vetor 2 - Digite o %dº valor: ", i+1)
		fmt.Scan(&v2[i])
	}

	// intercalação
	for i := 0; i < 10; i++ {
		v3[2*i] = v1[i]
		v3[2*i+1] = v2[i]
	}

	// impressão do resultado
	fmt.Println("\nVetor intercalado:")
	for i := 0; i < 20; i++ {
		fmt.Print(v3[i], " ")
	}
}
