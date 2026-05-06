package main

import "fmt"

func main() {
	var codigo int
	var vetor [10]float64

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])
	}

	fmt.Print("Digite o código (0, 1 ou 2): ")
	fmt.Scan(&codigo)

	if codigo == 0 {
		fmt.Println("Programa encerrado.")
		return
	} else if codigo == 1 {
		fmt.Println("Vetor na ordem direta:")
		for i := 0; i < 10; i++ {
			fmt.Printf("%.2f ", vetor[i])
		}
	} else if codigo == 2 {
		fmt.Println("Vetor na ordem inversa:")
		for i := 9; i >= 0; i-- {
			fmt.Printf("%.2f ", vetor[i])
		}
	} else {
		fmt.Println("Código inválido!")
	}
}
