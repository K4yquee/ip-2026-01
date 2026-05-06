package main

import "fmt"

func main() {
	var codigos [10]int
	var saldos [10]float64

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o código da conta %d: ", i+1)
		fmt.Scan(&codigos[i])

		for j := 0; j < i; j++ {
			if codigos[i] == codigos[j] {
				fmt.Println("Código já existente! Digite novamente.")
				i--
				break
			}
		}

		fmt.Printf("Digite o saldo da conta %d: ", i+1)
		fmt.Scan(&saldos[i])
	}

	opcao := 0

	for opcao != 4 {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Depósito")
		fmt.Println("2. Saque")
		fmt.Println("3. Ativo bancário")
		fmt.Println("4. Sair")
		fmt.Print("Escolha: ")
		fmt.Scan(&opcao)

		switch opcao {

		case 1:
			var codigo int
			var valor float64
			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&codigo)

			pos := -1

			for i := 0; i < 10; i++ {
				if codigos[i] == codigo {
					pos = i
					break
				}
			}

			if pos == -1 {
				fmt.Println("Conta não encontrada")
			} else {
				fmt.Print("Digite o valor do depósito: ")
				fmt.Scan(&valor)
				saldos[pos] += valor
				fmt.Println("Depósito realizado!")
			}

		case 2:
			var codigo int
			var valor float64
			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&codigo)

			pos := -1

			for i := 0; i < 10; i++ {
				if codigos[i] == codigo {
					pos = i
					break
				}
			}

			if pos == -1 {
				fmt.Println("Conta não encontrada")
			} else {
				fmt.Print("Digite o valor do saque: ")
				fmt.Scan(&valor)

				if saldos[pos] >= valor {
					saldos[pos] -= valor
					fmt.Println("Saque realizado!")
				} else {
					fmt.Println("Saldo insuficiente")
				}
			}

		case 3:
			total := 0.0

			for i := 0; i < 10; i++ {
				total += saldos[i]
			}

			fmt.Printf("Total no banco: %.2f\n", total)

		case 4:
			fmt.Println("Programa encerrado.")

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
