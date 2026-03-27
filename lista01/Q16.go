package main

import "fmt"

func main() {
	var sal, novosal float64

	fmt.Println("Digite o salario atual: ")
	fmt.Scan(&sal)

	if sal <= 300 {
		novosal = sal * 1.5
	} else {
		novosal = sal * 1.3
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", novosal)
}
