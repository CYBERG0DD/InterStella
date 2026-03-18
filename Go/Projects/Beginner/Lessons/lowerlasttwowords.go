package main

import (
	"fmt"
	"strings"
)

func lowerN(words []string, n int) []string {
	for char := range words {
		if char > len(words)-n {
			words[char] = strings.ToLower(words[char])
		}
	}
	return words
}

func main() {

	fmt.Println(lowerN([]string{"I", "LOVE", "YOU"}, 2))
}
