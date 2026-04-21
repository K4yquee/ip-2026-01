package main

import (
	"fmt"
	"math"
)

func main() {
	var S float64 = 0

	for i := 50; i >= 0; i-- {

		den := float64(2*i + 1)
		termo := 1.0 / (den * den * den)

		if i%2 == 0 {
			S += termo
		} else {
			S -= termo
		}
	}

	pi := math.Cbrt(S * 32)

	fmt.Printf("Valor de S: %.9f\n", S)
	fmt.Printf("Valor de pi ≈ %.9f\n", pi)
}
