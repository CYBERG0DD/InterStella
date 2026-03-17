package main

import (
	"fmt"
)

func check(word string) bool {
	tags := []string{"(up)", "(low)", "(cap)", "(hex)", "(bin)"}
	for _, tag := range tags {
		if word == tag {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(check("(up)"))
	fmt.Println(check("(low)"))
	fmt.Println(check("(cap)"))
	fmt.Println(check("(hex)"))
	fmt.Println(check("(bin)"))
	fmt.Println(check("(unknown)"))
}
