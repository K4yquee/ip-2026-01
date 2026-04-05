package main

import "fmt"

func main() {
   var idade int

   fmt.Println("Digite a sua idade: ")
   fmt.Scan(&idade)

   switch {
   case idade <= 2:
       fmt.Println("Você é um Recém-Nascido")
   case idade >= 3 && idade <= 11:
       fmt.Println("Você é uma Criança")
   case idade >= 12 && idade <= 19:
       fmt.Println("Você é um Adolescente")
   case idade >= 20 && idade <= 55:
       fmt.Println("Você é um adulto")
   case idade > 55:
       fmt.Println("Você é um idoso")
   }
}
