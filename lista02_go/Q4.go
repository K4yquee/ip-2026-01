package main

import (
   f "fmt"
   m "math"
)

func main() {
   var n, quadrado, raiz float64

   f.Println("Digite um numero: ")
   f.Scan(&n)

   quadrado = n * n
   raiz = (m.Sqrt(n))
   if n > 0 {
       f.Printf("A raiz quadrada do numero %.2f E %.2f\n", n, raiz)
   }
   if n < 0 {
       f.Println("O quadrado do numero ", n, "E", quadrado)
   }
}
