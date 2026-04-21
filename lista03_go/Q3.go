package main

import "fmt"

func main() {
	var joao, carlos, salj, salc float64
	var meses int

	fmt.Println("Digite o salário de carlos: ")
	fmt.Scan(&salc)

	salj = salc / 3

	carlos = salc
	joao = salj

	for joao < carlos {
		carlos = carlos * 1.02
		joao = joao * 1.05
		meses++
	}

	fmt.Printf("O salário de Carlos e %.2f, logo o de João e %.2f. \nDemorariam se %d meses ate joão ultrapassar carlos. \n Com João tendo %.2f e Carlos %.2f \n", salc, salj, meses, joao, carlos)
}
