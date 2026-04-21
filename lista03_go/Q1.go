package main

import "fmt"

func main() {
	var base, expo, resultado int

	fmt.Println("Digite a base da potência: ")
	fmt.Scan(&base)

	fmt.Println("Digite o expoente da potência: ")
	fmt.Scan(&expo)

	resultado = 1

	if base < 1 || expo < 1 {
		fmt.Println("Valores invalidos.")
		return
	}

	for i := 0; i < expo; i++ {
		resultado = resultado * base
	}

	fmt.Printf("O valor de %d elevado a %d e = %d", base, expo, resultado)
}
