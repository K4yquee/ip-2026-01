package main

import "fmt"

func main() {
	var n int
	var a1, a2, ai int

	fmt.Print("Digite o primeiro termo: ")
	fmt.Scan(&a1)

	fmt.Print("Digite o segundo termo: ")
	fmt.Scan(&a2)

	fmt.Print("Digite a quantidade de termos (N >= 3): ")
	fmt.Scan(&n)

	fmt.Printf("Sequência: %d %d ", a1, a2)

	for i := 3; i <= n; i++ {

		if i%2 != 0 {
			ai = a2 + a1
		} else {
			ai = a2 - a1
		}

		fmt.Printf("%d ", ai)

		a1 = a2
		a2 = ai
	}
}
