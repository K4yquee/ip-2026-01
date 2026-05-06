//A série de Fibonacci é formada pela sequência:
//1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
//Escreva um programa que armazene em um vetor os primeiros 50
//termos da série de Fibonacci. Após isso, o
//programa deve imprimir todos os valores armazenados.

package main

import "fmt"

func main() {
	var fib [50]int64

	// primeiros termos
	fib[0] = 1
	fib[1] = 1

	// calcula o restante
	for i := 2; i < 50; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	// imprime os valores
	for i := 0; i < 50; i++ {
		fmt.Println(fib[i])
	}
}

//VERIFICAR
