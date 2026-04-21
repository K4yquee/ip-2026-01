package main

import "fmt"

func main() {

	fmt.Println("Tabela do seno (aproximação por Maclaurin):\n")

	for i := 0; i <= 63; i++ {
		A := float64(i) / 10

		seno := A -
			(A*A*A)/6 +
			(A*A*A*A*A)/120 -
			(A*A*A*A*A*A*A)/5040

		fmt.Printf("A = %.1f -> sen(A) ≈ %.6f\n", A, seno)
	}
}
