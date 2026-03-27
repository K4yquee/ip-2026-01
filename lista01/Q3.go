package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Println("Digite o valor do primeiro numero: ")
	fmt.Scan(&n1)

	fmt.Println("Digite o valor do segundo numero: ")
	fmt.Scan(&n2)

	fmt.Println("Digite o valor do terceiro numero: ")
	fmt.Scan(&n3)

	if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
		fmt.Println("DIGITO INVALIDO")
	}

	numero := n1*100 + n2*10 + n3

	quadrado := numero * numero

	fmt.Printf("%d, %d\n", numero, quadrado)
}
