package main

import "fmt"

func main() {
   var n1, n2, soma, maior, menor int
   fmt.Println("Digite o primeiro numero: ")
   fmt.Scan(&n1)

   fmt.Println("Digite o segundo numero")
   fmt.Scan(&n2)

   soma = n1 + n2
   maior = soma + 8
   menor = soma - 5
   if soma > 20 {
       fmt.Println("O Resultado e:", soma, "+ 8 =", maior)
   }
   if soma < 20 {
       fmt.Println("O resultado e: ", soma, "- 5 =", menor)
   }
}
