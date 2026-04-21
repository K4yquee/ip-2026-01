package main

import "fmt"

func main() {
	var S float64 = 0

	for i := 0; i <= 14; i++ {
		numerador := 1 << i // 2^i
		denominador := (15 - i) * (15 - i)

		termo := float64(numerador) / float64(denominador)

		if i%2 == 0 {
			S += termo
		} else {
			S -= termo
		}
	}

	fmt.Printf("Valor de S: %.2f\n", S)
}
