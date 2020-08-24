package main

import "fmt"

func inc(x *int) {
	*x += 1
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	fibs := []int{0, 1, 1}

	return func() int {
		defer inc(&count)

		if count >= 0 && count <= 2 {
			return fibs[count]
		}

		// Extend the slice by 1 more element.
		fibs = append(fibs, 1)
		fibs[count] = fibs[count-1] + fibs[count-2]

		return fibs[count]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
