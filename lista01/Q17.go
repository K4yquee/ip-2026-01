package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("Digite o primeiro numero par: ")
	fmt.Scan(&x)

	fmt.Println("Digite tamanho da sequencia dele: ")
	fmt.Scan(&y)

	if x%2 == 0 {
		for i := 0; i < y; i++ {
			fmt.Printf("%d ", x)
			x += 2
		}
	} else {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	}
}
