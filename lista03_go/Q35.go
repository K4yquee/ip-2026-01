package main

import "fmt"

func main() {
	var n int
	var binario string

	binario = ""

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	for n > 0 {
		resto := n % 2
		binario = fmt.Sprint(resto) + binario
		n = n / 2
	}

	fmt.Println("Binário:", binario)
}
