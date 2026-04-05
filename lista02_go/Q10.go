package main

import "fmt"

func main() {
   var regiao, meio int

   fmt.Println("Informe o preço da passagem de acordo com a tabela a seguir")
   fmt.Println("1 - Região Norte R$ 500,00 (Ida e volta) | R$ 900,00 (Ida e Volta)")
   fmt.Println("2 - Região Nordeste R$ 350,00 (Ida) | R$ 650,00 (Ida e Volta)")
   fmt.Println("3 - Região Centro-Oeste R$ 350,00 (Ida) | R$ 600,00 (Ida e Volta)")
   fmt.Println("4 - Região Sul R$ 300,00 (Ida) | R$ 550,00 (Ida e Volta)")
  
   fmt.Println("Digite a região de destino (dentre os valores de 1 a 4): ")
   fmt.Scan(&regiao)

   fmt.Println("Digite se será somente ida (1) ou ida e volta (2): ")
   fmt.Scan(&meio)
  
   if regiao == 1 && meio == 1 {
       fmt.Println("O destino é a Região Norte como somente ida no valor de R$ - 500,00")
   } else if regiao == 1 && meio == 2 {
       fmt.Println("O destino é a Região Norte como ida e volta no valor de R$ - 900,00")
   }
   if regiao == 2 && meio == 1 {
       fmt.Println("O destino é a Região Nordeste como somente ida no valor de R$ - 350,00")
   } else if regiao == 2 && meio == 2 {
       fmt.Println("O destino é a Região Nordeste como ida e volta no valor de R$ - 650,00")
   }
   if regiao == 3 && meio == 1 {
       fmt.Println("O destino é a Região Centro-Oeste como somente ida no valor de R$ - 350,00")
   } else if regiao == 3 && meio == 2 {
       fmt.Println("O destino é a Região Centro-Oeste como ida e volta no valor de R$ - 600,00")
   }
   if regiao == 4 && meio == 1 {
       fmt.Println("O destino é a Região Sul como somente ida no valor de R$ - 300,00")
   } else if regiao == 4 && meio == 2 {
       fmt.Println("O destino é a Região Sul como ida e volta no valor de R$ - 550,00")
   }
  
}
