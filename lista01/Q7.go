package main

import "fmt"

func main() {
	var f, p float64

	fmt.Println("Digite o valor em fahrenheit: ")
	fmt.Scan(&f)

	fmt.Println("Digite o valor em polegadas: ")
	fmt.Scan(&p)

	c := (5*f - 160) / 9
	mm := p * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", c)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f mm\n", mm)
}
