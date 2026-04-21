package main

import "fmt"

func main() {
	var numero int
	var peso float64

	var maiorPeso, menorPeso float64
	var numGordo, numMagro int

	for i := 1; i <= 90; i++ {
		fmt.Printf("Digite o numero do boi %d: ", i)
		fmt.Scan(&numero)

		fmt.Printf("Digite o peso do boi %d: ", i)
		fmt.Scan(&peso)

		if i == 1 {
			maiorPeso = peso
			menorPeso = peso
			numGordo = numero
			numMagro = numero
		} else {
			if peso > maiorPeso {
				maiorPeso = peso
				numGordo = numero
			}

			if peso < menorPeso {
				menorPeso = peso
				numMagro = numero
			}
		}
	}

	fmt.Println("\nBoi mais gordo:")
	fmt.Println("Numero:", numGordo)
	fmt.Println("Peso:", maiorPeso)

	fmt.Println("\nBoi mais magro:")
	fmt.Println("Numero:", numMagro)
	fmt.Println("Peso:", menorPeso)
}
