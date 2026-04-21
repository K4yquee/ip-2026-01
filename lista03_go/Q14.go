package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Digite o valor de N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite o valor de N2: ")
	fmt.Scan(&n2)

	fmt.Println("Números primos no intervalo:")

	for i := n1; i <= n2; i++ {

		if i < 2 {
			continue
		}

		cont := 0

		for j := 1; j <= i; j++ {
			if i%j == 0 {
				cont++
			}
		}

		if cont == 2 {
			fmt.Println(i)
		}
	}
}
