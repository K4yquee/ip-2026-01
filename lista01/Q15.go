package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite um numero: ")
	fmt.Scan(&n)

	for i := 2; i <= n; i += 2 {
		q := i * i

		fmt.Printf("%d ao quadrado = %d\n", i, q)
	}
}
