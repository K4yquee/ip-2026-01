//Faça um programa que receba 100 valores numéricos, armazene-os em um
//vetor, calcule e imprima o valor do
//somatório dado a seguir:
//S=(b0-b99)
//3
//+ (b1-b98)
//3
//+ (b2-b97)
//3
//+ . . . + (b49-b50)
//3

package main

import "fmt"

func main() {
	var vetor [100]int
	var soma int

	// preenchimento automático
	for i := 0; i < 100; i++ {
		vetor[i] = i
	}

	// cálculo do somatório
	for i := 0; i < 50; i++ {
		diferenca := vetor[i] - vetor[99-i]
		soma += diferenca * diferenca * diferenca
	}

	fmt.Println("Resultado:", soma)
}

//VERIFICAR
