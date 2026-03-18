package main

import (
	"fmt"
	"strings"
)

func capitalizefirst(sentence string) string {
	if sentence == "" {
		return ""
	}
	return strings.ToUpper(sentence[:1]) + strings.ToLower(sentence[1:])
}

func main() {

	words := []string{"hello", "world", "from", "golang"}
	for _, word := range words {
		fmt.Println(capitalizefirst(word))
	}
}
