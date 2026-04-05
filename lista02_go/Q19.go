package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura, volume, area float64

	fmt.Println("Escolha a figura:")
	fmt.Println("1 - Cone")
	fmt.Println("2 - Cilindro")
	fmt.Println("3 - Esfera")
	fmt.Scan(&opcao)

	if opcao == 1 {
		fmt.Print("Digite o raio: ")
		fmt.Scan(&raio)

		fmt.Print("Digite a altura: ")
		fmt.Scan(&altura)

		volume = (math.Pi * raio * raio * altura) / 3
		area = math.Pi * raio * math.Sqrt(raio*raio+altura*altura)

		fmt.Println("Cone:")
		fmt.Printf("Volume: %.2f\n", volume)
		fmt.Printf("Área: %.2f\n", area)

	} else if opcao == 2 {
		fmt.Print("Digite o raio: ")
		fmt.Scan(&raio)

		fmt.Print("Digite a altura: ")
		fmt.Scan(&altura)

		volume = math.Pi * raio * raio * altura
		area = 2 * math.Pi * raio * altura

		fmt.Println("Cilindro:")
		fmt.Printf("Volume: %.2f\n", volume)
		fmt.Printf("Área: %.2f\n", area)

	} else if opcao == 3 {
		fmt.Print("Digite o raio: ")
		fmt.Scan(&raio)

		volume = (4.0 / 3.0) * math.Pi * math.Pow(raio, 3)
		area = 4 * math.Pi * raio * raio

		fmt.Println("Esfera:")
		fmt.Printf("Volume: %.2f\n", volume)
		fmt.Printf("Área: %.2f\n", area)

	} else {
		fmt.Println("Opção inválida")
	}
}
