package main

import (
	"fmt"
	"math"
)

// func Sqrt(x float64) float64 {
// 	z := 1.0

// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Println(z)
// 	}

// 	return z
// }

func Sqrt(x float64) float64 {
	z := 1.0
	oldZ := 0.0

	for {
		oldZ = z
		z -= (z*z - x) / (2 * z)

		if math.Abs(oldZ-z) < 0.000000000001 {
			return z
		}

		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(" => ", Sqrt(3), "\tExpected: ", math.Sqrt(3))
}
