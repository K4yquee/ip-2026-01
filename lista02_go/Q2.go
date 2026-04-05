package main

import "fmt"

func main() {
   var n int

   fmt.Println("Digite um numero: ")
   fmt.Scan(&n)

   if n > 0 {
       fmt.Println("O numero  positivo")
   }
   if n < 0 {
       fmt.Println("O numero e negativo")
   }
   if n == 0 {
       fmt.Println("O numero e nulo")
   }
}
