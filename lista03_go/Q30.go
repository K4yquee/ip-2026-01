package main

import (
	"fmt"
	"math"
)

func main() {

	for R := 0.0; R <= 20.0; R += 0.5 {

		volume := (4.0 / 3.0) * math.Pi * R * R * R

		fmt.Printf("R = %.1f -> Volume = %.2f\n", R, volume)
	}
}
