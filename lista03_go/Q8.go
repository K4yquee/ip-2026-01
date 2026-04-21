package main

import "fmt"

func main() {
   var soma int

   o := 0
   soma = 0
   for i := 0; i <= 20; i++ {
       soma += i
       o = i
   }

   fmt.Printf("%d, a soma dos números é: %d\n", o, soma)

}

Q8
package main

import "fmt"

func main() {
   var soma int

   soma = 0
   for i := 0; i <= 20; i++ {
       soma += i
       fmt.Println(i)
   }

   fmt.Printf("A soma dos números é: %d\n", soma)

}
