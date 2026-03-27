package main

import f "fmt"

func main() {
	var n int

	f.Println("Digite o numero: ")
	f.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		f.Println("O NUMERO E DIVISIVEL")
	} else {
		f.Println("O NUMERO NAO E DIVISIVEL")
	}
}
