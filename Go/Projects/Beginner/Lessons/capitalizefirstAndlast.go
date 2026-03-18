package main

import (
	"fmt"
	"strings"
)

func capitalizeLastFirstLetter(sentence string) string {
	if len(sentence) == 0 {
		return sentence
	}
	return strings.ToUpper(sentence[:1]) + strings.ToLower(sentence[1:len(sentence)-1]) + strings.ToUpper(string(sentence[len(sentence)-1]))
}

func main() {

	words := []string{
		"hello", "world", "from", "golang", "Sweet", "To", "Code",
	}
	for _, char := range words {
		fmt.Println(capitalizeLastFirstLetter(char))
	}
}
