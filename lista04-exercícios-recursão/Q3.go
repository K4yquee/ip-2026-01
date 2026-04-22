package main

import "fmt"

func inverter(valores []int, inicio, fim int) {
	if inicio >= fim {
		return
	}

	valores[inicio], valores[fim] = valores[fim], valores[inicio]

	inverter(valores, inicio+1, fim-1)
}

func main() {
	valores := []int{1, 2, 3, 4, 5}

	inverter(valores, 0, len(valores)-1)

	fmt.Println("Array invertido:", valores)
}
