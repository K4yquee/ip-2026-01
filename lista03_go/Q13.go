package main

import "fmt"

func main() {
	var h float64

	for i := 1; i <= 50; i++ {
		h += float64(2*i-1) / float64(i)
	}

	fmt.Printf("Valor de H: %.2f\n", h)
}
