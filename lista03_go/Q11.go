package main

import "fmt"

func main() {
	var n int
	var fatorial int

	fatorial = 1

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Valor inválido! Não existe fatorial de número negativo.")
		return
	}

	for i := 1; i <= n; i++ {
		fatorial *= i
	}

	fmt.Printf("O fatorial de %d é: %d\n", n, fatorial)
}
