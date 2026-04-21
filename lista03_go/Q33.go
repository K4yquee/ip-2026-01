package main

import "fmt"

func main() {
	var n1, n2 int
	var quociente int = 0

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	for n1 >= n2 {
		n1 -= n2
		quociente++
	}

	fmt.Printf("Quociente = %d\n", quociente)
	fmt.Printf("Resto = %d\n", n1)
}
