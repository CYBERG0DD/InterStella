package main

import (
	"fmt"
	"strings"
)

func CapitalizeLastLetter(sentence string) string {
	if sentence == "" {
		return ""
	}
	return strings.ToLower(sentence[:len(sentence)-1]) + strings.ToUpper(string(sentence[len(sentence)-1]))
}

func main() {

	words := []string{"hello", "world", "from", "golang", "Sweet", "To", "Code"}
	for _, char := range words {
		fmt.Println(CapitalizeLastLetter(char))
	}
}
