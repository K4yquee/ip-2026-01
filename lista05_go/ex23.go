package main

import "fmt"

func main() {
	var janela [24]int
	var corredor [24]int

	opcao := -1

	for opcao != 0 {
		fmt.Println("\n--- VENDA DE PASSAGENS ---")
		fmt.Println("1. Janela")
		fmt.Println("2. Corredor")
		fmt.Println("0. Sair")
		fmt.Print("Escolha: ")
		fmt.Scan(&opcao)

		if opcao == 0 {
			fmt.Println("Encerrando programa.")
			break
		}

		var vetor *[24]int
		tipo := ""

		if opcao == 1 {
			vetor = &janela
			tipo = "Janela"
		} else if opcao == 2 {
			vetor = &corredor
			tipo = "Corredor"
		} else {
			fmt.Println("Opção inválida!")
			continue
		}

		fmt.Printf("\nPoltronas disponíveis (%s):\n", tipo)
		disponivel := false

		for i := 0; i < 24; i++ {
			if vetor[i] == 0 {
				fmt.Printf("%d ", i)
				disponivel = true
			}
		}

		if !disponivel {
			fmt.Println("\nNão há poltronas disponíveis nessa opção!")
			continue
		}

		var escolha int
		fmt.Print("\nEscolha o número da poltrona: ")
		fmt.Scan(&escolha)

		if escolha < 0 || escolha >= 24 {
			fmt.Println("Poltrona inválida!")
			continue
		}

		if vetor[escolha] == 1 {
			fmt.Println("Poltrona já ocupada!")
		} else {
			vetor[escolha] = 1
			fmt.Println("Poltrona reservada com sucesso!")
		}

		cheio := true

		for i := 0; i < 24; i++ {
			if janela[i] == 0 || corredor[i] == 0 {
				cheio = false
				break
			}
		}

		if cheio {
			fmt.Println("Ônibus completamente cheio!")
			break
		}
	}
}
