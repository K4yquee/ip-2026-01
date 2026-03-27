package main

import "fmt"

func main() {
	var n int
	var soma float64 = 0

	fmt.Println("Digite um numero inteiro positivo: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}

	for k := 1; k <= n; k++ {
		soma += 1.0 / float64(k)
	}

	fmt.Printf("%.6f\n", soma)
}
