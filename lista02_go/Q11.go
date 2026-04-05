package main

import "fmt"

func main() {
   var x, f int

   fmt.Println("Digite o valor de x: ")
   fmt.Scan(&x)

   f = 8 / (2 - x)

   fmt.Printf("O valor de f(x) = 8 / (2 - x) é: %d\n", f)
}
