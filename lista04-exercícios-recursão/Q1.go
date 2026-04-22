package main

import "fmt"

func potencia(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * potencia(x, n-1)
}

func main() {
	var x, n int

	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)

	fmt.Print("Digite o valor de n: ")
	fmt.Scan(&n)

	resultado := potencia(x, n)

	fmt.Println("Resultado:", resultado)
}
