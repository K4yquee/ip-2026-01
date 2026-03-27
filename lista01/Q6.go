package main

import "fmt"

func main() {
	var n int
	var f, c float64

	fmt.Println("Digite a quantidade de testes: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		fmt.Println("Digite o valor em Fahrenheit: ")
		fmt.Scan(&f)

		c = (5 * (f - 32)) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", f, c)
	}
}
