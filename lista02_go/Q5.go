package main

import "fmt"

func main() {
   var n int
   fmt.Println("Digite um numero: ")
   fmt.Scan(&n)

   if n%5 == 0 {
       fmt.Printf("O numero %d e divisivel por 5", n)
   } else {
       fmt.Printf("O numero %d nao e divisivel por 5 \n", n)
   }

}
