package main

import "fmt"

func main() {

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {

			if j > i {
				fmt.Printf("(%d,%d) ", i, j)
			}
		}
		fmt.Println()
	}
}
