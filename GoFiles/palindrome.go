package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The Logic: Building the machine to reverse and compare
func isPalindrome(text string) bool {
	runes := []rune(text)
	n := len(runes)

	reversed := make([]rune, n)

	// Fill the reversed slice
	for i := 0; i < n; i++ {
		reversed[i] = runes[n-1-i]
	}

	// Compare original vs reversed
	for i := 0; i < n; i++ {
		if runes[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter value: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		value := strings.TrimSpace(input)

		if value == "" {
			fmt.Println("Error: Value cannot be empty")
			continue
		}

		if isPalindrome(value) {
			fmt.Printf("'%s' is a palindrome! ✅\n", value)
		} else {
			fmt.Printf("'%s' is NOT a palindrome. ❌\n", value)
		}

		break // Exit loop after a successful check
	}
}