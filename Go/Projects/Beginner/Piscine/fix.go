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
				vowels := "aeiouAEIOUhH"
				if strings.ContainsAny(string(nextword[0]), vowels) {
					if word == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {

	fmt.Println(fixAToAn("a honest mistake"))
}
