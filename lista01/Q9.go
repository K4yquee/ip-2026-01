package main

import f "fmt"

func main() {
	var a, b, c float64

	f.Println("Digite o valor de a: ")
	f.Scan(&a)

	f.Println("Digite o valor de b: ")
	f.Scan(&b)

	f.Println("Digite o valor de c: ")
	f.Scan(&c)

	delta := b*b - 4*a*c

	f.Printf("O VALOR DE DELTA E = %.2f\n", delta)

}
