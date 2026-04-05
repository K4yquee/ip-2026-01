package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número inteiro de 3 casas: ")
	fmt.Scan(&num)

	if num >= 100 && num <= 999 {
		dezena := (num / 10) % 10
		fmt.Println("Algarismo da casa das dezenas:", dezena)
	} else {
		fmt.Println("Erro: o número não tem 3 casas.")
	}
}
