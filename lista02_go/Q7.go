package main

import "fmt"

func main() {
	var a, b, c, menor, maior, inter int

	fmt.Println("Digite o valor inteiro (A): ")
	fmt.Scan(&a)

	fmt.Println("Digite o valor inteiro (B): ")
	fmt.Scan(&b)

	fmt.Println("Digite o valor inteiro (C): ")
	fmt.Scan(&c)

	maior = a
	menor = a

	if b < menor {
		menor = b
	}
	if b > maior {
		maior = b
	}

	if c < menor {
		menor = c
	}
	if c > maior {
		maior = c
	}

	inter = a + b + c - menor - maior

	fmt.Printf("Ordem crescente: %d, %d, %d\n", menor, inter, maior)
}
