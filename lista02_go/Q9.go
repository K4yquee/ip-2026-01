package main

import "fmt"

func main() {
	var valor, lucro float64

	fmt.Println("Digite o valor da compra: ")
	fmt.Scan(&valor)

	if valor < 10 {
		lucro = valor * 1.7
	}

	if valor >= 10 && valor < 30 {
		lucro = valor * 1.5
	}

	if valor >= 30 && valor < 50 {
		lucro = valor * 1.4
	}

	if valor >= 50 {
		lucro = valor * 1.3
	}

	fmt.Printf("O valor da compra e %.2f e o valor da venda e %.2f\n", valor, lucro)
}
