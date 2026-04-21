package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	mmc := n1

	for {
		if mmc%n1 == 0 && mmc%n2 == 0 {
			break
		}
		mmc++
	}

	fmt.Printf("MMC = %d\n", mmc)
}
