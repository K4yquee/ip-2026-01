package main

import "fmt"

func main() {
	var num [10]int
	var divis [5]int

	fmt.Println("Digite os 10 números do vetor principal:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&num[i])
	}

	fmt.Println("Digite os 5 números divisores:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Divisor %d: ", i+1)
		fmt.Scan(&divis[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("\nNúmero %d:\n", num[i])

		encontrou := false

		for j := 0; j < 5; j++ {
			if divis[j] != 0 && num[i]%divis[j] == 0 {
				fmt.Printf("Divisível por %d na posição %d\n", divis[j], j)
				encontrou = true
			}
		}

		if !encontrou {
			fmt.Println("Não é divisível por nenhum número do segundo vetor")
		}
	}
}
