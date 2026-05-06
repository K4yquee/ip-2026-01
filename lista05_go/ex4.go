// Faça um programa que leia um vetor A de 10 posições, contendo números inteiros. Determine e mostre, a
// seguir, quais elementos de A estão repetidos e quantas vezes cada um se repete.
package main

import "fmt"

func main() {
	var a [10]int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d número: ", i+1)
		fmt.Scan(&a[i])
	}

	for i := 0; i < 10; i++ {
		jaApareceu := false

		for k := 0; k < i; k++ {
			if a[k] == a[i] {
				jaApareceu = true
			}
		}

		if jaApareceu {
			continue
		}

		contador := 0
		for j := 0; j < 10; j++ {
			if a[i] == a[j] {
				contador++
			}
		}

		// mostra só se repetir
		if contador > 1 {
			fmt.Printf("O número %d aparece %d vezes\n", a[i], contador)
		}
	}
}
