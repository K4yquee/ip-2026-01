package main

import "fmt"

func main() {
	var a1, r, n, termo, soma int
	soma = 0

	fmt.Println("Digite o termo a1: ")
	fmt.Scan(&a1)

	fmt.Println("Digite a razão: ")
	fmt.Scan(&r)

	fmt.Println("Digite o numero de repetiçoes: ")
	fmt.Scan(&n)

	termo = a1

	for i := 0; i < n; i++ {
		soma += termo
		termo += r
	}
	fmt.Println(soma)
}
