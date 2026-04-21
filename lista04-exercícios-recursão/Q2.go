package main

import "fmt"

func main() {
	valores := []float64{2.5, 3.0, 4.5, 1.0}
	soma := 0.0

	for i := 0; i < len(valores); i++ {
		soma += valores[i]
	}

	fmt.Println("Soma:", soma)
}
