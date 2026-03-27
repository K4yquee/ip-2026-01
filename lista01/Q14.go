package main

import (
	f "fmt"
	m "math"
)

func main() {
	var h, a float64

	f.Println("Digite a altura da piramide: ")
	f.Scan(&h)

	f.Println("Digite o lado do hexagono: ")
	f.Scan(&a)

	ab := (3 * a * a * m.Sqrt(3)) / 2

	v := (ab * h) / 3

	f.Printf("O VOLUME DA PIRAMIDE E = %.2f\n", v)
}
