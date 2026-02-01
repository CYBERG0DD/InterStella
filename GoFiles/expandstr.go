package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	words := strings.Fields(input)
	if len(words) == 0 {
		return
	}
	result := strings.Join(words, "   ")
	fmt.Println(result)
}
