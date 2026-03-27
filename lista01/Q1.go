package main

import "fmt"

func main() {
	var n1, n2, n3, media float64

	fmt.Println("Digite o valor da primeira nota: ")
	fmt.Scan(&n1)

	fmt.Println("Digite o valor da segunda nota: ")
	fmt.Scan(&n2)

	fmt.Println("Digite o valor da terceira nota: ")
	fmt.Scan(&n3)

	media = (n1 + n2 + n3) / 3

	fmt.Printf("MEDIA = %.2f\n", media)

	if media >= 6 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
