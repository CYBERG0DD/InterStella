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
	var err error
	var input string
	var choice int
	var history []string

	const (
		Red    = "\033[31m"
		Blue   = "\033[34m"
		Yellow = "\033[33m"
		Green  = "\033[32m"
		Reset  = "\033[0m"
	)

	for {

		fmt.Println()
		fmt.Printf("%s============ [OPERATION GOPHER'S PROTOCOL] ============%s\n", Blue, Reset)
		fmt.Println()

		listProtocols := []string{"UPPER CASE", "LOWER CASE", "CAPITALIZE FIRST LETTER", "TITLE CASE", "SNAKE CASE", "REVERSE", "EXIT"}
		for i, protocol := range listProtocols {
			fmt.Printf("%s[%d]. %s%s\n", Green, i+1, protocol, Reset)
		}

		fmt.Println()
		fmt.Printf("%sSelect a protocol:(1-7): %s", Yellow, Reset)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%sError reading input: %v%s\n", Red, err, Reset)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput cannot be empty. Please try again.%s\n", Red, Yellow, Reset)
			continue
		}

		choice, err = strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 7 {
			fmt.Printf("\n%s[INPUT ERROR] ---->> %sOops looks like you have gone out of range from the outlisted options. Try again.%s\n", Red, Yellow, Reset)
			continue
		}

		if choice == 7 {
			fmt.Printf("\n%sAre you sure you want to exit? (yes/no): %s", Red, Reset)
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Printf("%sError reading input: %v%s\n", Red, err, Reset)
				continue
			}

			input = strings.TrimSpace(strings.ToLower(input))
			if input == "" {
				fmt.Printf("\n%s[INPUT ERROR] ---->> Please fill in this field for a successful exit.%s\n", Red, Reset)
				continue
			}

			switch input {
			case "yes", "y":
				fmt.Println()
				fmt.Printf("%s============ [SESSION HISTORY] ============%s\n", Blue, Reset)
				fmt.Println()

				if len(history) == 0 {
					fmt.Printf("%sNo operations were performed in this session.%s\n", Yellow, Reset)
				} else {
					for i, entry := range history {
						fmt.Printf("%s[%d] %s%s\n", Green, i+1, entry, Reset)
					}
				}

				fmt.Println()
				fmt.Printf("%s=================================================%s\n", Blue, Reset)
				fmt.Printf("\n%sExiting... GOPHER'S PROTOCOL!%s\n\n", Green, Reset)
				return

			case "no", "n":
				fmt.Printf("\n%sExit cancelled. Returning to the main menu...%s\n", Yellow, Reset)
				continue

			default:
				fmt.Printf("\n%s[INPUT ERROR] ---->> Invalid input. Please enter 'yes/y' or 'no/n'.%s\n", Red, Reset)
				continue
			}
		}

		for {

			if choice == 1 {
				fmt.Printf("\n%sEnter a string to convert to %s[UPPER CASE]: %s", Yellow, Green, Reset)
				input, err = reader.ReadString('\n')
				if err != nil {
					fmt.Printf("%sError reading input: %v%s\n", Red, err, Reset)
					continue
				}

				input = strings.TrimSpace(input)
				if input == "" {
					fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput cannot be empty. Please try again.%s\n", Red, Yellow, Reset)
					continue
				}
				isValid := true
				for _, char := range input {
					if unicode.IsDigit(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
						fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput contains invalid characters that are not letters. Please try again.%s\n", Red, Yellow, Reset)
						isValid = false
						break
					}
				}
				if !isValid {
					continue
				}
				result := strings.ToUpper(input)
				fmt.Println()
				fmt.Printf("%s==============[CONVERTED TO UPPER CASE]==============%s\n", Green, Reset)
				fmt.Printf("%sResult: %s%s%s\n", Green, Yellow, result, Reset)
				fmt.Printf("%sYou converted: %s'%s' to: '%s'%s\n", Green, Yellow, input, result, Reset)
				fmt.Printf("%s=====================================================%s\n", Green, Reset)
				fmt.Println()
				history = append(history, fmt.Sprintf("UPPER CASE   | '%s' --> '%s'", input, result))
				break
			}

			if choice == 2 {
				fmt.Printf("\n%sEnter a string to convert to %s[LOWER CASE]: %s", Yellow, Green, Reset)
				input, err = reader.ReadString('\n')
				if err != nil {
					fmt.Printf("%sError reading input: %v%s\n", Red, err, Reset)
					continue
				}

				input = strings.TrimSpace(input)
				if input == "" {
					fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput cannot be empty. Please try again.%s\n", Red, Yellow, Reset)
					continue
				}
				isValid := true
				for _, char := range input {
					if unicode.IsDigit(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
						fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput contains invalid characters that are not letters. Please try again.%s\n", Red, Yellow, Reset)
						isValid = false
						break
					}
				}
				if !isValid {
					continue
				}
				result := strings.ToLower(input)
				fmt.Println()
				fmt.Printf("%s==============[CONVERTED TO LOWER CASE]==============%s\n", Green, Reset)
				fmt.Printf("%sResult: %s%s%s\n", Green, Yellow, result, Reset)
				fmt.Printf("%sYou converted: %s'%s' to: '%s'%s\n", Green, Yellow, input, result, Reset)
				fmt.Printf("%s=====================================================%s\n", Green, Reset)
				fmt.Println()
				history = append(history, fmt.Sprintf("LOWER CASE   | '%s' --> '%s'", input, result))
				break
			}

			if choice == 3 {
				fmt.Printf("\n%sEnter a string to convert to %s[CAPITALIZE FIRST LETTER]: %s", Yellow, Green, Reset)
				input, err = reader.ReadString('\n')
				if err != nil {
					fmt.Printf("%sError reading input: %v%s\n", Red, err, Reset)
					continue
				}

				input = strings.TrimSpace(input)
				if input == "" {
					fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput cannot be empty. Please try again.%s\n", Red, Yellow, Reset)
					continue
				}
				isValid := true
				for _, char := range input {
					if unicode.IsDigit(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
						fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput contains invalid characters that are not letters. Please try again.%s\n", Red, Yellow, Reset)
						isValid = false
						break
					}
				}
				if !isValid {
					continue
				}
				result := strings.ToUpper(input[:1]) + strings.ToLower(input[1:])
				fmt.Println()
				fmt.Printf("%s==============[CONVERTED TO CAPITALIZE FIRST LETTER]==============%s\n", Green, Reset)
				fmt.Printf("%sResult: %s%s%s\n", Green, Yellow, result, Reset)
				fmt.Printf("%sYou converted: %s'%s' to: '%s'%s\n", Green, Yellow, input, result, Reset)
				fmt.Printf("%s=================================================================%s\n", Green, Reset)
				fmt.Println()
				history = append(history, fmt.Sprintf("CAPITALIZE   | '%s' --> '%s'", input, result))
				break
			}

			if choice == 4 {
				fmt.Printf("\n%sEnter a string to convert to %s[TITLE CASE]: %s", Yellow, Green, Reset)
				input, err = reader.ReadString('\n')
				if err != nil {
					fmt.Printf("%sError reading input: %v%s\n", Red, err, Reset)
					continue
				}

				input = strings.TrimSpace(input)
				if input == "" {
					fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput cannot be empty. Please try again.%s\n", Red, Yellow, Reset)
					continue
				}
				isValid := true
				for _, char := range input {
					if unicode.IsDigit(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
						fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput contains invalid characters that are not letters. Please try again.%s\n", Red, Yellow, Reset)
						isValid = false
						break
					}
				}
				if !isValid {
					continue
				}
				result := strings.Fields(input)
				for i, char := range result {
					if len(char) > 0 {
						result[i] = strings.ToUpper(char[:1]) + strings.ToLower(char[1:])
					}
				}
				finalResult := strings.Join(result, " ")
				fmt.Println()
				fmt.Printf("%s==============[CONVERTED TO TITLE CASE]==============%s\n", Green, Reset)
				fmt.Printf("%sResult: %s%s%s\n", Green, Yellow, finalResult, Reset)
				fmt.Printf("%sYou converted: %s'%s' to: '%s'%s\n", Green, Yellow, input, finalResult, Reset)
				fmt.Printf("%s=====================================================%s\n", Green, Reset)
				fmt.Println()
				history = append(history, fmt.Sprintf("TITLE CASE   | '%s' --> '%s'", input, finalResult))
				break
			}

			if choice == 5 {
				fmt.Printf("\n%sEnter a string to convert to %s[SNAKE CASE]: %s", Yellow, Green, Reset)
				input, err = reader.ReadString('\n')
				if err != nil {
					fmt.Printf("%sError reading input: %v%s\n", Red, err, Reset)
					continue
				}

				input = strings.TrimSpace(input)
				if input == "" {
					fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput cannot be empty. Please try again.%s\n", Red, Yellow, Reset)
					continue
				}
				isValid := true
				for _, char := range input {
					if unicode.IsDigit(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
						fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput contains invalid characters that are not letters. Please try again.%s\n", Red, Yellow, Reset)
						isValid = false
						break
					}
				}
				if !isValid {
					continue
				}
				result := strings.ReplaceAll(input, " ", "_")
				fmt.Println()
				fmt.Printf("%s==============[CONVERTED TO SNAKE CASE]==============%s\n", Green, Reset)
				fmt.Printf("%sResult: %s%s%s\n", Green, Yellow, result, Reset)
				fmt.Printf("%sYou converted: %s'%s' to: '%s'%s\n", Green, Yellow, input, result, Reset)
				fmt.Printf("%s=====================================================%s\n", Green, Reset)
				fmt.Println()
				history = append(history, fmt.Sprintf("SNAKE CASE   | '%s' --> '%s'", input, result))
				break
			}

			if choice == 6 {
				fmt.Printf("\n%sEnter a string to convert to %s[REVERSE]: %s", Yellow, Green, Reset)
				input, err = reader.ReadString('\n')
				if err != nil {
					fmt.Printf("%sError reading input: %v%s\n", Red, err, Reset)
					continue
				}

				input = strings.TrimSpace(input)
				if input == "" {
					fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput cannot be empty. Please try again.%s\n", Red, Yellow, Reset)
					continue
				}
				isValid := true
				for _, char := range input {
					if unicode.IsDigit(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
						fmt.Printf("\n%s[INPUT ERROR] ---->> %sInput contains invalid characters that are not letters. Please try again.%s\n", Red, Yellow, Reset)
						isValid = false
						break
					}
				}
				if !isValid {
					continue
				}
				runes := []rune(input)
				for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
					runes[i], runes[j] = runes[j], runes[i]
				}
				finalResult := string(runes)
				fmt.Println()
				fmt.Printf("%s==============[CONVERTED TO REVERSE]==============%s\n", Green, Reset)
				fmt.Printf("%sResult: %s%s%s\n", Green, Yellow, finalResult, Reset)
				fmt.Printf("%sYou converted: %s'%s' to: '%s'%s\n", Green, Yellow, input, finalResult, Reset)
				fmt.Printf("%s==================================================%s\n", Green, Reset)
				fmt.Println()
				history = append(history, fmt.Sprintf("REVERSE      | '%s' --> '%s'", input, finalResult))
				break
			}
		}
	}
}
