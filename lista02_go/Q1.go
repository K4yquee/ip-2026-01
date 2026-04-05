package main

import "fmt"

func main() {
   var n int
   fmt.Println("Escreva um numero: ")
   fmt.Scan(&n)

   if n%2 == 0 {
       fmt.Println("O numero e par")
   } else {
       fmt.Println("O numero e impar")
   }
}
