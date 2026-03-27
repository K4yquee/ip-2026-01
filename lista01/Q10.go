package main

import f "fmt"

func main() {
	var a, b, c, d float64

	f.Println("Digite o valor de a: ")
	f.Scan(&a)

	f.Println("Digite o valor de b: ")
	f.Scan(&b)

	f.Println("Digite o valor de c: ")
	f.Scan(&c)

	f.Println("Digite o valor de d: ")
	f.Scan(&d)

	determinante := (a * d) - (b * c)

	f.Printf("O VALOR DO DETERMINANTE E = %.2f\n", determinante)

}
