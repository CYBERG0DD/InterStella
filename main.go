package main

import "github.com/01-edu/z01"

func main() {
	for bid := '0'; bid <= '9'; bid++ {
		z01.PrintRune(bid)
	}
	z01.PrintRune('\n')
}
