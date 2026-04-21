package main

import "fmt"

func main() {
	valores := []int{1, 2, 3, 4, 5}

	n := len(valores)

	for i := 0; i < n/2; i++ {
		valores[i], valores[n-1-i] = valores[n-1-i], valores[i]
	}

	fmt.Println("Array invertido:", valores)
}
