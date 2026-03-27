package main

import "fmt"

func main() {
	var salario, consumo float64

	fmt.Println("Digite o valor do salario minimo: ")
	fmt.Scan(&salario)

	fmt.Println("Digite a quantia consumida em KW: ")
	fmt.Scan(&consumo)

	valorKW := (0.7 * salario) / 100
	custo := valorKW * consumo
	custoDesconto := custo * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", valorKW)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}
