package main

import (
	"fmt"
	"unicode"
)

func ISpunc(word string) string {
	for _, i := range word {
		if unicode.IsPunct(i) {
			return "It is a Punctuation"
		}
	}
	return "Not a punctuation"
}

func main() {
	fmt.Println(ISpunc("Hi, my name is Apex your new AI assistant!"))
	fmt.Println(ISpunc("Hey are you ready for vacation"))
}
