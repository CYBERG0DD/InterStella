package main

import (
	"fmt"
	"regexp"
)

func fixarticles(word string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(word, "'$1'")
}

func main() {

	fmt.Println(fixarticles("This is a ' test ' string with ' extra spaces ' around the words."))
}
