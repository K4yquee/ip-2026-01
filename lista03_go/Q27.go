package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var S float64 = 0

	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)

	fatorial := 1.0
	potencia := 1.0

	for i := 0; i < 20; i++ {

		if i > 0 {
			potencia *= x * x
			fatorial *= float64(2*i-1) * float64(2*i)
		}

		termo := potencia / fatorial

		if i%2 == 0 {
			S += termo
		} else {
			S -= termo
		}
	}

	cosReal := math.Cos(x)

	diferenca := S - cosReal

	fmt.Printf("a) cos(%.2f) pela série ≈ %.6f\n", x, S)
	fmt.Printf("b) diferença ≈ %.6f\n", diferenca)
	fmt.Printf("c) cos(%.2f) real ≈ %.6f\n", x, cosReal)
}
