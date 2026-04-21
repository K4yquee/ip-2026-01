package main

import "fmt"

func main() {
	var altura, peso, alt10e20 float64
	var idade, maior50, totalpessoas, entre10e20, pmenor40kg int

	continuar := 1
	for continuar == 1 {
		fmt.Println("Digite a idade da pessoa: ")
		fmt.Scan(&idade)

		fmt.Println("Digite a altura da pessoa: ")
		fmt.Scan(&altura)

		fmt.Println("Digite o peso da pessoa: ")
		fmt.Scan(&peso)

		totalpessoas++

		if idade > 50 {
			maior50++
		}
		if peso < 40 {
			pmenor40kg++
		}
		if idade >= 10 && idade <= 20 {
			alt10e20 += altura
			entre10e20++
		}

		fmt.Println("Deseja continuar? (1 - Sim - Outro - Não)")
		fmt.Scan(&continuar)
	}

	media10e20 := alt10e20 / float64(entre10e20)
	pormenor40kg := (float64(pmenor40kg) / float64(totalpessoas)) * 100

	fmt.Printf("A quantida de pessoas maiores de 50 é: %d\n", maior50)
	fmt.Printf("A média das alturas das pessoas com idade entre 10 e 20 anos é: %.2f\n", media10e20)
	fmt.Printf("A porcentagem de pessoas com peso inferior a 40kg entre todas as analisadas é: %.2f\n", pormenor40kg)
}
