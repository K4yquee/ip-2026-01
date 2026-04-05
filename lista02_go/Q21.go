package main

import "fmt"

func main() {
	var id int
	var n1, n2, n3, me, mf float64
	var conceito string

	fmt.Print("Digite o número do aluno: ")
	fmt.Scan(&id)

	fmt.Print("Digite a nota 1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite a nota 2: ")
	fmt.Scan(&n2)

	fmt.Print("Digite a nota 3: ")
	fmt.Scan(&n3)

	fmt.Print("Digite a média dos exercícios: ")
	fmt.Scan(&me)

	mf = (n1 + n2*2 + n3*3 + me) / 7

	if mf >= 9 {
		conceito = "A"
	} else if mf >= 7.5 {
		conceito = "B"
	} else if mf >= 6 {
		conceito = "C"
	} else if mf >= 4 {
		conceito = "D"
	} else {
		conceito = "E"
	}

	fmt.Println("\nAluno:", id)
	fmt.Println("Notas:", n1, n2, n3)
	fmt.Println("Média dos exercícios:", me)
	fmt.Printf("Média final: %.2f\n", mf)
	fmt.Println("Conceito:", conceito)

	if conceito == "A" || conceito == "B" || conceito == "C" {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
