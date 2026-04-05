package main

import "fmt"

func main() {
	var preco, valor float64
	var codigo int

	fmt.Print("Digite o preço do produto: ")
	fmt.Scan(&preco)

	fmt.Print("Digite o código de pagamento (1 a 4): ")
	fmt.Scan(&codigo)

	if codigo == 1 {
		valor = preco * 0.9
	} else if codigo == 2 {
		valor = preco * 0.95
	} else if codigo == 3 {
		valor = preco
	} else if codigo == 4 {
		valor = preco * 1.1
	} else {
		fmt.Println("Código inválido!")
		return
	}

	fmt.Printf("Valor a pagar: R$ %.2f\n", valor)
}
