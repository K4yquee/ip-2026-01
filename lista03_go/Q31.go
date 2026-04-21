package main

import "fmt"

func main() {
	var grao uint64 = 1
	var total uint64 = 0

	for i := 0; i < 64; i++ {
		total += grao
		grao *= 2
	}

	fmt.Println("Total de grãos:", total)
}
