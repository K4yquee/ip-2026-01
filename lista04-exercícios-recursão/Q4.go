package main

import "fmt"

func decimalParaBinario(n int) string {
	if n == 0 {
		return ""
	}

	resto := n % 2
	return decimalParaBinario(n/2) + fmt.Sprint(resto)
}

func main() {
	var n int

	fmt.Print("Digite um número decimal: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Binário: 0")
	} else {
		binario := decimalParaBinario(n)
		fmt.Println("Binário:", binario)
	}
}
