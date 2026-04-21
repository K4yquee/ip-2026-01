package main

import "fmt"

func main() {
	var num int

	for {
		fmt.Print("Digite um número (<=0 para parar): ")
		fmt.Scan(&num)

		if num <= 0 {
			break
		}

		ehQuadrado := false

		for i := 1; i*i <= num; i++ {
			if i*i == num {
				ehQuadrado = true
				break
			}
		}

		if ehQuadrado {
			fmt.Println("É quadrado perfeito")
		} else {
			fmt.Println("Não é quadrado perfeito")
		}
	}
}
