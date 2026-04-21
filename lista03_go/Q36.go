package main

import "fmt"

func main() {
	var n int
	var hexa string

	hexa = ""

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Hexadecimal: 0")
		return
	}

	for n > 0 {
		resto := n % 16

		if resto < 10 {
			hexa = fmt.Sprint(resto) + hexa
		} else {
			hexa = string('A'+(resto-10)) + hexa
		}

		n = n / 16
	}

	fmt.Println("Hexadecimal:", hexa)
}
