package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite a quantidade de testes: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var total int
		var pPop, pGer, pArq, pCad float64

		fmt.Println("Digite a quantia total de pessoas: ")
		fmt.Scan(&total)

		fmt.Println("Digite a porcentagem de pessoas na categoria popular: ")
		fmt.Scan(&pPop)

		fmt.Println("Digite a porcentagem de pessoas na categoria geral: ")
		fmt.Scan(&pGer)

		fmt.Println("Digite a porcentagem de pessoas na categoria arquibancada: ")
		fmt.Scan(&pArq)

		fmt.Println("Digite a porcentagem de pessoas na categoria cadeiras: ")
		fmt.Scan(&pCad)

		pop := float64(total) * (pPop / 100)
		ger := float64(total) * (pGer / 100)
		arq := float64(total) * (pArq / 100)
		cad := float64(total) * (pCad / 100)

		renda := pop*1 + ger*5 + arq*10 + cad*20

		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}
