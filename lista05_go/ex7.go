package main

import "fmt"

func main() {
	var vetor []int // slice (dinâmico)

	for i := 100; i >= 1; i-- {
		if i%2 != 0 {
			vetor = append(vetor, i)
		}
	}

	for _, valor := range vetor {
		fmt.Println(valor)
	}
}
