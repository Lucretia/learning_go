package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for i := range result {
		result[i] = make([]uint8, dx)
	}

	for y := 0; y < len(result[0]); y++ {
		for x := 0; x < len(result[y]); x++ {
			//			result[y][x] = uint8((x+y)/2)
			//			result[y][x] = uint8(x * y)
			result[y][x] = uint8(x ^ y)
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
