package main

import (
	"fmt"
	"strings"
)

func CapFirsttLetter(sentence string) string {
	if sentence == "" {
		return ""
	}
	return strings.ToUpper(sentence[:1]) + strings.ToLower(sentence[1:])
}

func main() {

	words := []string{"hello", "world", "from", "golang", "Sweet", "To", "Code"}
	for _, i := range words {
		fmt.Println(CapFirsttLetter(i))
	}
}
