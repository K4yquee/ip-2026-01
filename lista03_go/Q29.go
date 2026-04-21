package main

import "fmt"

func main() {
	var N int
	var S int = 0

	fmt.Print("Digite um valor para N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		S += i
	}

	fmt.Printf("Soma de 1 até %d = %d\n", N, S)
}
