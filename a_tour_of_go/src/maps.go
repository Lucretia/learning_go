package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func FindWords(word string, w []string) int {
	var result = 0

	for i := range w {
		if word == w[i] {
			result++
		}
	}

	return result
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int, len(words))

	// Initialise the map
	for i := range words {
		result[words[i]] = FindWords(words[i], words)
	}

	return result
	//	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
