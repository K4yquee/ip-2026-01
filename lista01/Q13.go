package main

import f "fmt"

func main() {
	var nota float64

	f.Println("Digite a sua nota: ")
	f.Scan(&nota)

	if nota >= 9.0 {
		f.Printf("NOTA = %.1f CONCEITO = A\n", nota)
	} else if nota >= 7.5 {
		f.Printf("Nota = %.1f CONCEITO = B\n", nota)
	} else if nota >= 6.0 {
		f.Printf("Nota = %.1f CONCEITO = C\n", nota)
	} else {
		f.Printf("Nota = %.1f CONCEITO = D\n", nota)
	}
}
