package main

import "fmt"

func somaRecursiva(valores []float64, indice int) float64 {
	if indice == len(valores) {
		return 0
	}
	return valores[indice] + somaRecursiva(valores, indice+1)
}

func main() {
	valores := []float64{2.5, 3.0, 4.5, 1.0}

	soma := somaRecursiva(valores, 0)

	fmt.Println("Soma:", soma)
}
