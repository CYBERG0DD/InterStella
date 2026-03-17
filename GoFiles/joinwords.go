package main

import (
	"fmt"
	"strings"
)

func joinwords(a, b string) string {

	word := []string{a, b}
	result := strings.Join(word, "")
	return result
}

func main() {

	fmt.Println(joinwords("hello", "world"))
}
