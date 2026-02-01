package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str := os.Args[1]
	from := os.Args[2]
	to := os.Args[3]
	found := false
	for _, ch := range str {
		if string(ch) == from {
			found = true
			break
		}
	}
	if !found {
		fmt.Println(str)
		return
	}
	result := ""
	for _, ch := range str {
		if string(ch) == from {
			result += to
		} else {
			result += string(ch)
		}
	}
	fmt.Println(result)
}
