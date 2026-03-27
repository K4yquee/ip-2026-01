package main

import f "fmt"

func main() {
	var horas int
	var total float64

	f.Println("Defina o total de horas: ")
	f.Scan(&horas)

	blocos := horas / 3
	resto := horas % 3

	total = float64(blocos*10 + resto*5)

	f.Printf("O VALOR A PAGAR E %.2f\n", total)
}
