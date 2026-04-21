package main

import "fmt"

func main() {
	var S float64 = 0
	fatorial := 1.0 
	for i := 0; i < 20; i++ {

		if i > 0 {
			fatorial *= float64(i)
		}

		termo := float64(100-i) / fatorial
		S += termo
	}

	fmt.Printf("Soma: %.2f\n", S)
}
