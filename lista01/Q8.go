package main

import "fmt"

func main() {
	var r, h float64
	const pi = 3.14159
	const custoMetro = 100.0

	fmt.Println("Escreva o raio da lata: ")
	fmt.Scan(&r)

	fmt.Println("Escreva a altura da lata: ")
	fmt.Scan(&h)

	area := 2*pi*r*r + 2*pi*r*h

	custo := area * custoMetro

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
