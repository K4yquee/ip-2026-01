package main

import "fmt"

func main() {
   var numero, soma, quantidade, quantidadepares, maior, menor, pares, quantidadeimpares int
   var media, porcentagemimpar, mediapares float64

   continuar := 1

   for continuar == 1 {

       fmt.Println("Digite um número: ")
       fmt.Scan(&numero)
       soma += numero
       quantidade++

       media = float64(soma) / float64(quantidade)

       if numero%2 == 0 {
           pares = pares + numero
           quantidadepares++
       }

       if numero%2 != 0 {
           quantidadeimpares++
       }

       if quantidade == 1 {
           maior = numero
           menor = numero
       } else {
           if numero > maior {
               maior = numero
           }
           if numero < menor {
               menor = numero
           }
       }
       porcentagemimpar = float64(quantidadeimpares) / float64(quantidade) * 100
       mediapares = float64(pares) / float64(quantidadepares)

       fmt.Println("Deseja adicionar mais números? (30000 - não. 1 - sim): ")
       fmt.Scan(&continuar)

       if continuar == 30000 {
           break
       } else {
           continuar = 1
       }
   }

   fmt.Printf("A soma dos números digitados é: %d\n", soma)
   fmt.Printf("A quantidade de números digitados é: %d\n", quantidade)
   fmt.Printf("A média dos números digitados é: %.2f\n", media)
   fmt.Printf("O maior número digitado é: %d\n", maior)
   fmt.Printf("O menor número digitado é: %d\n", menor)
   fmt.Printf("A média dos números pares é: %.2f\n", mediapares)
   fmt.Printf("A porcentagem de números impares entre todos os números digitados é: %.2f\n", porcentagemimpar)
}
