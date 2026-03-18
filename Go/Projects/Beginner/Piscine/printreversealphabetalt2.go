package piscine

import (
	"fmt"
)

func main() {
	// Predefined alphabet slices without casting
	upper := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	lower := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z',
	}

	// First line: reverse order from 'Z' to 'a', alternating uppercase/lowercase
	for i := 25; i >= 0; i-- {
		if (25-i)%2 == 0 {
			fmt.Printf("%c", upper[i])
		} else {
			fmt.Printf("%c", lower[i])
		}
	}
	fmt.Println()

	for i := 0; i < 26; i++ {
		if i%2 == 0 {
			fmt.Printf("%c", upper[i])
		} else {
			fmt.Printf("%c", lower[i])
		}
	}
	fmt.Println()
}
