package main

import (
	"fmt"
	"math"
)

type SqrtError struct {
	msg string
}

func (e *SqrtError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &SqrtError{"Cannot pass Sqrt a value less than zero!"}
	}

	z := x / 2
	oldZ := 0.0

	for {
		oldZ = z
		z -= (z*z - x) / (2 * z)

		if math.Abs(oldZ-z) < 0.000000000001 {
			return z, nil
		}

		// fmt.Println(z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
