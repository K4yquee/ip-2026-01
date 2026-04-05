package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("Digite A: ")
	fmt.Scan(&a)

	fmt.Print("Digite B: ")
	fmt.Scan(&b)

	fmt.Print("Digite C: ")
	fmt.Scan(&c)

	if a == 0 {
		fmt.Println("Não é equação do segundo grau.")
		return
	}

	delta := b*b - 4*a*c

	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Println("RAIZ ÚNICA")
		fmt.Println("x =", x)
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)

		fmt.Println("RAÍZES DISTINTAS")
		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
	}
}
