package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []int

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i := 0; i < len(digits); i++ {
		for k := i + 1; k < len(digits); k++ {
			if digits[i] > digits[k] {
				digits[i], digits[k] = digits[k], digits[i]
			}
		}
	}
	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
}
