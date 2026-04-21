package main

import "fmt"

func main() {
   var numero int

   fmt.Println("Digite um numero triangular: ")
   fmt.Scan(&numero)

   o := 0

   for i := 0; i*(i+1)*(i+2) < numero; {
       i++
       if i*(i+1)*(i+2) == numero {
           fmt.Printf("%d é um número triangular. Dos numeros %d, %d, %d\n", numero, i, i+1, i+2)
       }
       o = i
   }
   if o*(o+1)*(o+2) != numero {
       fmt.Printf("%d não é um número triangular. \n", numero)
   }
}
