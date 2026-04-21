package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite a quantidade de termos: ")
	fmt.Scan(&n)

	a := 0
	b := 1

	fmt.Printf("Sequência de Fibonacci: %d - %d", a, b)

	for i := 2; i < n; i++ {
		proximo := a + b
		fmt.Printf(" - %d", proximo)

		a = b
		b = proximo
	}
}
