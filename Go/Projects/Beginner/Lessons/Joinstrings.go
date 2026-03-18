package main

import (
	"fmt"
	"strings"
)

func JoinWords(a, b string) string {

	words := []string{a, b}
	result := strings.Join(words, " ")
	return result
}

func main() {

	fmt.Println(JoinWords("Hello", "World"))
}
