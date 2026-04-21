package main

import "fmt"

func main() {
	var x, n int
	resultado := 1

	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)

	fmt.Print("Digite o valor de n: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		resultado *= x
	}

	fmt.Println("Resultado:", resultado)
}
