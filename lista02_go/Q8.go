package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite um numero maior que 20 e menor que 90")
	fmt.Scan(&n)

	if n > 20 && n < 90 {
		fmt.Printf("Sucesso, o numero %d esta dentro da faixa", n)
	} else {
		fmt.Printf("Incorreto, o numero %d nao esta dentro da faixa de valores", n)
	}
}
