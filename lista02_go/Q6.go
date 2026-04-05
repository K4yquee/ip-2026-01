package main

import "fmt"

func main() {
   var a, b int

   fmt.Println("Digite um numero (a): ")
   fmt.Scan(&a)

   fmt.Println("Digite um numero (b): ")
   fmt.Scan(&b)

   if a%b == 0 {
       fmt.Printf("O numero %d e divisivel pelo numero %d\n", a, b)
   } else {
       fmt.Printf("O numero %d nao e divisivel pelo numero %d\n", a, b)
   }
}
