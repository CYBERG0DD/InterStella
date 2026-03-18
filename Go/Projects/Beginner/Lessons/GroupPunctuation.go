package main

import (
	"fmt"
	"strings"
)

func isGroupPunctuation(s string) bool {
	punctuations := "!@#$%^&*()_+-=~`[]{}|;:',.<>/?"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(punctuations, char) {
			count++
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(isGroupPunctuation("..?,"))
	fmt.Println(isGroupPunctuation(", , $"))
	fmt.Println(isGroupPunctuation(".hello?"))
}
