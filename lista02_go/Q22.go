package main

import "fmt"

func main() {
	var matricula int
	var horasExtra float64

	const salarioMinimo = 788.0
	const valorHoraExtra = 10.0

	var salarioExtra, salarioBruto float64
	var inss, ir, salarioLiquido float64

	fmt.Print("Digite a matrícula: ")
	fmt.Scan(&matricula)

	fmt.Print("Digite a quantidade de horas extras: ")
	fmt.Scan(&horasExtra)

	salarioExtra = horasExtra * valorHoraExtra
	salarioBruto = 3*salarioMinimo + salarioExtra

	if salarioBruto > 1500 {
		inss = salarioBruto * 0.12
	} else {
		inss = 0
	}

	ir = salarioBruto * 0.20

	salarioLiquido = salarioBruto - inss - ir

	fmt.Println("\nMatrícula:", matricula)
	fmt.Printf("Salário bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Desconto INSS: R$ %.2f\n", inss)
	fmt.Printf("Desconto IR: R$ %.2f\n", ir)
	fmt.Printf("Salário líquido: R$ %.2f\n", salarioLiquido)
}
