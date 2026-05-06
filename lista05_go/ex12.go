//Faça um programa que:
//- leia um conjunto de valores inteiros correspondentes a 15 notas
//dos alunos de uma turma. As notas variam de
//0 a 10;
//- calcule a frequência absoluta e a frequência relativa de cada nota;
//- imprima uma tabela contendo os valores das notas (0 a 10) e suas respectivas
//frequências absoluta e relativa.
//Observações:
//- Frequência absoluta de um nota é o número de vezes em que ela aparece no conjunto de dados;
//- Frequência relativa é a frequência absoluta dividida pelo número total de dados.

package main

import "fmt"

func main() {
	var nota int
	var freq [11]int
	total := 15

	// leitura das notas
	for i := 0; i < total; i++ {
		fmt.Printf("Digite a %dª nota (0 a 10): ", i+1)
		fmt.Scan(&nota)

		if nota >= 0 && nota <= 10 {
			freq[nota]++
		} else {
			fmt.Println("Nota inválida, tente novamente.")
			i-- // repete a leitura
		}
	}

	// impressão da tabela
	fmt.Println("\nNota | Freq. Absoluta | Freq. Relativa")
	for i := 0; i <= 10; i++ {
		relativa := float64(freq[i]) / float64(total)
		fmt.Printf("%4d | %15d | %15.2f\n", i, freq[i], relativa)
	}
}

//VERIFICAR
