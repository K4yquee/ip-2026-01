package main

import "fmt"

func main() {
	var b, n int
	var resultado int

	resultado = 1

	fmt.Print("Digite o valor da base (b): ")
	fmt.Scan(&b)

	fmt.Print("Digite o valor do expoente (n): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		resultado *= b
	}

	fmt.Printf("Resultado: %d\n", resultado)
}
