package piscine

import "fmt"

func PrintMemory(arr [10]byte) {
	// Print hex dump: 2 digits + space per byte
	for i := 0; i < 10; i++ {
		b := arr[i]
		// High nibble
		if b/16 < 10 {
			putchar('0' + b/16)
		} else {
			putchar('a' + (b/16 - 10))
		}
		// Low nibble
		if b%16 < 10 {
			putchar('0' + b%16)
		} else {
			putchar('a' + (b%16 - 10))
		}
		putchar(' ')
	}
	putchar(' ')

	// Print ASCII: printable chars or dot
	for i := 0; i < 10; i++ {
		b := arr[i]
		if b >= 32 && b <= 126 {
			putchar(b)
		} else {
			putchar('.')
		}
	}
	putchar('\n')
}

// Print single char (requires fmt import in main)
func putchar(c byte) {
	fmt.Printf("%c", c)
}
