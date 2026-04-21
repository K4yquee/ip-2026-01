package main

import "fmt"

func main() {
	var n1, n2, media, mclass float64
	var alunos, aprovados, exame, reprovados int

	fmt.Println("Você ira avaliar a média aritmética de quantos alunos?: ")
	fmt.Scan(&alunos)

	for i := 0; i < alunos; i++ {
		fmt.Printf("Digite a primeira nota do aluno número %d: ", i+1)
		fmt.Scan(&n1)
		fmt.Printf("Digite a segunda nota do aluno número %d: ", i+1)
		fmt.Scan(&n2)

		media = (n1 + n2) / 2

		mclass += media
		if media >= 7 {
			aprovados++
		} else if media >= 3 && media < 7 {
			exame++
		} else if media < 3 {
			reprovados++
		}
	}

	mclass = mclass / float64(alunos)

	fmt.Printf("%d alunos foram aprovados.\n", aprovados)
	fmt.Printf("%d alunos faram exame.\n", exame)
	fmt.Printf("%d alunos foram reprovados.\n", reprovados)
	fmt.Printf("A média da classe é: %.2f", mclass)
}
