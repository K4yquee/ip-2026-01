package main

import "fmt"

func main() {
	var s float64

	for i := 1; i <= 37; i++ {
		num := 38 - i
		s += float64(num*(num+1)) / float64(i)
	}

	fmt.Printf("Valor de S: %.2f\n", s)
}
