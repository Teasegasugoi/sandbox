package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		pre_z := z
		z -= (z*z - x) / (2 * z)
		if abs := math.Abs(z - pre_z); abs < 1e-10 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
