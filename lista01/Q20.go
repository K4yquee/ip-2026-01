package main

import "fmt"

func main() {
	var h, m, s int

	fmt.Println("Digite a quantia de horas: ")
	fmt.Scan(&h)

	fmt.Println("Digite a quantia de minutos: ")
	fmt.Scan(&m)

	fmt.Println("Digite a quantia de segundos: ")
	fmt.Scan(&s)

	total := h*3600 + m*60 + s

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", total)
}
