package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // To read full line input
	var name string
	var withdrawalMethod string
	var amount int

	fmt.Println("POCKET OPTION FINANCIAL SECTION")

	// --- Get user's name ---
	for {
		fmt.Print("Enter full name: ")
		nameInput, _ := reader.ReadString('\n')
		name = strings.TrimSpace(nameInput)

		if name == "" {
			fmt.Println("Name cannot be empty. Please try again.")
			continue
		}

		if _, err := strconv.Atoi(name); err == nil {
			fmt.Println("Name cannot be a number. Please enter a valid name.")
			continue
		}

		break
	}

	fmt.Printf("Welcome %s, kindly select your payment method below.\n", name)

	// --- Select withdrawal method ---
	for {
		fmt.Println("1. Bank Transfer")
		fmt.Println("2. Others")
		fmt.Print("Enter withdrawal method (1 or 2): ")
		input, _ := reader.ReadString('\n')
		withdrawalMethod = strings.TrimSpace(input)
		if withdrawalMethod == "1" || withdrawalMethod == "2" {
			break
		}
		fmt.Println("Invalid option. Please enter 1 or 2.")
	}

	// --- BANK TRANSFER ---
	if withdrawalMethod == "1" {
		banks := []string{"Opay", "Palmpay", "Kuda", "Moniepoint"}
		var bankChoice string

		fmt.Println("Choose your bank:")
		for i, bank := range banks {
			fmt.Printf("%d. %s\n", i+1, bank)
		}

		for {
			fmt.Print("Enter bank option (1-4): ")
			input, _ := reader.ReadString('\n')
			bankChoice = strings.TrimSpace(input)
			if bankChoice == "1" || bankChoice == "2" || bankChoice == "3" || bankChoice == "4" {
				break
			}
			fmt.Println("Invalid bank selection. Please try again.")
		}

		for {
			fmt.Print("Enter amount to withdraw: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			num, err := strconv.Atoi(input)
			if err != nil || num <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			amount = num
			break
		}

		fmt.Printf("The sum of %d has been withdrawn successfully from %s!\n", amount, banks[toIndex(bankChoice)])
		return
	}

	// --- OTHERS (Cryptos etc.) ---
	if withdrawalMethod == "2" {
		others := []string{"BTC", "Ethereum", "TON Wallet", "Binance", "Crypto"}
		var otherChoice string
		var wallet string

		fmt.Println("Choose option:")
		for i, option := range others {
			fmt.Printf("%d. %s\n", i+1, option)
		}

		for {
			fmt.Print("Enter option (1-5): ")
			input, _ := reader.ReadString('\n')
			otherChoice = strings.TrimSpace(input)
			if otherChoice == "1" || otherChoice == "2" || otherChoice == "3" || otherChoice == "4" || otherChoice == "5" {
				break
			}
			fmt.Println("Invalid option. Please try again.")
		}

		for {
			fmt.Printf("Enter a valid %s wallet address: ", others[toIndex(otherChoice)])
			walletInput, _ := reader.ReadString('\n')
			wallet = strings.TrimSpace(walletInput)
			if wallet != "" {
				break
			}
			fmt.Println("Wallet address cannot be empty. Please try again.")
		}

		for {
			fmt.Print("Enter amount to withdraw: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			num, err := strconv.Atoi(input)
			if err != nil || num <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			amount = num
			break
		}

		fmt.Printf("The sum of %d has been withdrawn successfully to %s wallet (%s). Funds will arrive in 3 business days.\n",
			amount, others[toIndex(otherChoice)], wallet)
		return
	}
}

// Convert string option "1"-"5" to index 0-4
func toIndex(option string) int {
	idx, _ := strconv.Atoi(option)
	return idx - 1
}
