package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// --- STEP 1: INITIALIZATION ---
	score := 0
	trials := 0
	maxTrials := 10
	// In Go, we use a map[string]bool to act like a 'set'
	usedWords := make(map[string]bool)

	fmt.Println("--- Welcome to the Palindrome Quest! ---")
	fmt.Print("Would you love to play? (yes/no): ")

	ask, _ := reader.ReadString('\n')
	if strings.TrimSpace(strings.ToLower(ask)) != "yes" {
		fmt.Println("Maybe next time!")
		return
	}

	fmt.Print("Enter your username: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("\nOkay %s, you have %d attempts. Go!\n", name, maxTrials)

	// --- STEP 2: THE GAME LOOP ---
	for trials < maxTrials {
		fmt.Printf("\n[%d/%d] Enter keyword: ", trials+1, maxTrials)
		input, _ := reader.ReadString('\n')
		word := strings.TrimSpace(strings.ToLower(input))

		// 1. Check if empty
		if word == "" {
			fmt.Println("Don't waste a trial on an empty string!")
			continue
		}

		// 2. Check for duplicates (The Map Logic)
		if usedWords[word] {
			fmt.Printf("You already used '%s'! Think of a new one.\n", word)
			continue
		}

		// 3. Process the word
		trials++
		usedWords[word] = true // Mark as used

		if isPalindrome(word) {
			fmt.Printf("✅ Yes! '%s' is a palindrome.\n", word)
			score++
		} else {
			fmt.Printf("❌ No! '%s' is not a palindrome.\n", word)
		}

		// 4. Warning Logic
		if trials == 7 {
			fmt.Println("⚠️ WARNING: Only 3 trials left!")
		}
	}

	// --- STEP 3: FINAL SCORE ---
	fmt.Println("\n" + strings.Repeat("=", 20))
	fmt.Printf("GAME OVER, %s!\n", name)
	fmt.Printf("Total Palindromes Found: %d\n", score)

	if score >= 8 {
		fmt.Println("🏆 Great job! Well done!")
	} else {
		fmt.Println("📉 You did not reach the expected end. Keep practicing!")
	}
}

// Helper function for the core logic
func isPalindrome(s string) bool {
	// Remove spaces
	clean := strings.ReplaceAll(s, " ", "")
	runes := []rune(clean)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}
