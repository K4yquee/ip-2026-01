package main

import "fmt"

func main() {
	var n int
	var decimal int = 0
	var base int = 1

	fmt.Print("Digite um número em base 8: ")
	fmt.Scan(&n)

	for n > 0 {
		digito := n % 10

		if digito > 7 {
			fmt.Println("Número inválido para base 8!")
			return
		}

		decimal += digito * base
		base *= 8
		n = n / 10
	}

	fmt.Println("Em base 10:", decimal)
}
