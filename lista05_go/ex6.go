package main

import "fmt"

func main() {
   var vetor [100]int

   a := 0

   for i := 100; i >= 1; i-- {
       vetor[a] = i
       a++
   }

   for i := 0; i < 100; i++ {
       fmt.Println(vetor[i])
   }
}

q7
package main

import "fmt"

func main() {
   var vetor [100]int

   a := 0
   for i := 100; i >= 1; i-- {
       if i%2 != 0 {
           vetor[a] = i
           a++
       }
   }

   for i := 0; i < 100; i++ {
       fmt.Println(vetor[i])
   }

}