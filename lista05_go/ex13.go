//Uma grande firma deseja saber quais os três empregados mais recentes. Faça um programa para ler um número
//indeterminado de informações (máximo de 100), contendo o número do empregado e o número de meses de
//trabalho deste empregado. Imprima os três empregados que entraram para trabalhar mais recentemente na
//firma. Obs.: a última informação contém os dois números iguais a zero. Assuma que não existem dois ou mais
//empregados admitidos no mesmo mês.

package main

import "fmt"

func main() {
	var codigo, meses int

	// inicializa com valores bem altos
	m1, m2, m3 := 999999, 999999, 999999
	c1, c2, c3 := 0, 0, 0

	for {
		fmt.Print("Digite código e meses (0 0 para parar): ")
		fmt.Scan(&codigo, &meses)

		if codigo == 0 && meses == 0 {
			break
		}

		if meses < m1 {
			m3, c3 = m2, c2
			m2, c2 = m1, c1
			m1, c1 = meses, codigo

		} else if meses < m2 {
			m3, c3 = m2, c2
			m2, c2 = meses, codigo

		} else if meses < m3 {
			m3, c3 = meses, codigo
		}
	}

	fmt.Println("\nTrês empregados mais recentes:")
	fmt.Printf("1º: Código %d - %d meses\n", c1, m1)
	fmt.Printf("2º: Código %d - %d meses\n", c2, m2)
	fmt.Printf("3º: Código %d - %d meses\n", c3, m3)
}
