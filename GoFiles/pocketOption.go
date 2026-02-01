package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("POCKET OPTION FINANCIAL SECTION")

	// Get user's name - TRAPS if only numbers entered
	name := getValidName(reader)

	fmt.Printf("Welcome %s, kindly select your payment method below.\n", name)

	// Select withdrawal method
	withdrawalMethod := getWithdrawalMethod(reader)

	if withdrawalMethod == "1" {
		handleBankTransfer(reader)
	} else {
		handleOthers(reader)
	}
}

func getValidName(reader *bufio.Reader) string {
	for {
		fmt.Print("Enter full name: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		name := strings.TrimSpace(input)

		// Trap user in loop if input is ONLY numbers (no letters)
		if isOnlyNumbers(name) {
			fmt.Println("Invalid name! Names must not contain numbers, Please try again.")
			continue
		}

		if name != "" {
			return name
		}
		fmt.Println("Name cannot be empty. Please try again.")
	}
}

// Check if string contains ONLY digits (traps numeric-only input)
func isOnlyNumbers(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0 // Only return true if non-empty string of digits
}

func getWithdrawalMethod(reader *bufio.Reader) string {
	options := map[string]bool{"1": true, "2": true}
	for {
		fmt.Println("1. Bank Transfer")
		fmt.Println("2. Others")
		fmt.Print("Enter withdrawal method (1 or 2): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		choice := strings.TrimSpace(input)
		if options[choice] {
			return choice
		}
		fmt.Println("Invalid option. Please enter 1 or 2.")
	}
}

func handleBankTransfer(reader *bufio.Reader) {
	banks := []string{"Opay", "Palmpay", "Kuda", "Moniepoint"}

	fmt.Println("Choose your bank:")
	for i, bank := range banks {
		fmt.Printf("%d. %s\n", i+1, bank)
	}

	bankChoice := getValidChoice(reader, "Enter bank option (1-4): ", 1, 4)
	bankIndex := toIndex(bankChoice)

	amount := getValidAmount(reader, "Enter amount to withdraw: ")

	fmt.Printf("The sum of $%d has been withdrawn successfully from %s!\n",
		amount, banks[bankIndex])
}

func handleOthers(reader *bufio.Reader) {
	others := []string{"BTC", "Ethereum", "TON Wallet", "Binance", "Crypto"}

	fmt.Println("Choose option:")
	for i, option := range others {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	otherChoice := getValidChoice(reader, "Enter option (1-5): ", 1, 5)
	otherIndex := toIndex(otherChoice)

	wallet := getNonEmptyInput(reader,
		fmt.Sprintf("Enter a valid %s wallet address: ", others[otherIndex]),
		"Wallet address cannot be empty. Please try again.")

	amount := getValidAmount(reader, "Enter amount to withdraw: ")

	fmt.Printf("The sum of $%d has been withdrawn successfully to %s wallet (%s). Funds will arrive in 3 business days.\n",
		amount, others[otherIndex], wallet)
}

func getNonEmptyInput(reader *bufio.Reader, prompt, errorMsg string) string {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		result := strings.TrimSpace(input)
		if result != "" {
			return result
		}
		fmt.Println(errorMsg)
	}
}

func getValidChoice(reader *bufio.Reader, prompt string, min, max int) string {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		choice := strings.TrimSpace(input)
		if num, err := strconv.Atoi(choice); err == nil && num >= min && num <= max {
			return choice
		}
		fmt.Printf("Invalid selection. Please enter a number between %d-%d.\n", min, max)
	}
}

func getValidAmount(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		amountStr := strings.TrimSpace(input)
		amount, err := strconv.Atoi(amountStr)
		if err != nil || amount <= 0 {
			fmt.Println("Invalid amount. Please enter a positive number.")
			continue
		}
		return amount
	}
}

func toIndex(option string) int {
	idx, err := strconv.Atoi(option)
	if err != nil {
		return 0 // Safe fallback
	}
	return idx - 1
}
