package main

import (
	"fmt"
	"unicode"
)

func IsPunctuation(word string) bool {
	for _, i := range word {
		if unicode.IsPunct(i) {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(IsPunctuation("Hello, World!."))
	fmt.Println(IsPunctuation("Hello World"))
}
