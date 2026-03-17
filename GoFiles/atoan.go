package main

import (
	"fmt"
	"strings"
)

func fixAToAn(text string) string {
	words := strings.Split(text, " ")
	for i, word := range words {
		if word == "a" || word == "A" {
			if i+1 < len(words) {
				nextword := words[i+1]
				vowels := "aeiouAEIOU"
				if strings.ContainsAny(string(nextword[0]), vowels) {
					if word == "A" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {

	fmt.Println(fixAToAn("There it was. A amazing rock. A honest man. A book"))
}
