package main

import "fmt"

func main() {
	var n1, n2, multiplicacao int

	fmt.Println("Digite o valor de N1: ")
	fmt.Scan(&n1)

	fmt.Println("Digite o valor de N2: ")
	fmt.Scan(&n2)

	for i := 0; i < n2; i++ {
		multiplicacao += n1
	}

	fmt.Printf("O resultado é: %d", multiplicacao)
}
