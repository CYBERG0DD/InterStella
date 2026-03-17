package main

import (
	"fmt"
)

func firstletter(word string) string {
	if len(word) > 0 {
		return string(word[1])
	}
	return word
}

func main() {

	fmt.Println(firstletter("Cyber_G0dd"))
}
