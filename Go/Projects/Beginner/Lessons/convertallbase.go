package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	const (
		Red    = "\033[31m"
		Blue   = "\033[32m"
		Yellow = "\033[33m"
		Green  = "\033[34m"
		Reset  = "\033[0m"
	)

	for {

		var choice int
		var err error
		var input string

		fmt.Println()
		//("================= BASE CONVERTER =================\n")
		for {
			fmt.Println(Red + "\n====== [MENU SECTION] ======" + Reset)
			fmt.Println()

			slice := []string{
				"Hexadecimal to All",
				"Binary to All",
				"Decimal to All",
				"Exit converter",
			}

			for i, list := range slice {
				fmt.Printf("  %s[%d].  %s%s\n", Green, i+1, list, Reset)
				// fmt.Println()
			}

			fmt.Printf("%s\nSelect an option: %s", Yellow, Reset)
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

			choice, err = strconv.Atoi(input)
			if err != nil || choice < 1 || choice > 4 {
				fmt.Printf("%s\n[RANGE ERROR] ---->> Oops, looks like you have surpassed the outlisted options. Try again%s%s\n", Red, Yellow, Reset)
				continue
			}

			if choice == 4 {
				fmt.Println()
				fmt.Printf("%sAre you really sure you want to exit: %s", Red, Reset)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if input == "" {
					fmt.Printf("%s\n[ERROR] ---->> %sPlease fill in this field for a successful exit!%s\n\n", Red, Yellow, Reset)
					continue
				}
				switch strings.ToLower(input) {
				case "yes", "y":
					fmt.Printf("%s\nExiting converter........%s\n", Red, Reset)
					return

				case "no", "n":
					continue

				default:
					fmt.Println()
					fmt.Printf("%s[Invalid input]. %sPlease type 'yes/y' or 'no/n' for a successful exit.%s\n\n", Red, Yellow, Reset)
				}
			}

			break
		}

		//("================= [HEXADECIMAL SECTION]=================")
		if choice == 1 {
			for {
				fmt.Printf("%s\n------ CONVERT FROM HEXADECIMAL TO ANY BASE ------%s\n", Green, Reset)

				fmt.Printf("%sEnter HEX value: %s", Yellow, Reset)
				hexInput, _ := reader.ReadString('\n')
				hexInput = strings.TrimSpace(hexInput)

				result, err := strconv.ParseInt(hexInput, 16, 64)
				if err != nil {
					fmt.Printf("%s\n[INVALID HEX CODE] ---->> %s'%s' is not a valid [HEXADECIMAL] code%s\n\n", Red, Yellow, hexInput, Reset)
					continue
				}

				fmt.Printf("%sEnter base number: %s", Yellow, Reset)
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)

				targetBase, err := strconv.Atoi(input)
				if err != nil {
					fmt.Printf("%s\n[ERROR] ---->> Please enter a valid number%s\n", Red, Reset)
					continue
				}

				if targetBase != 2 && targetBase != 10 && targetBase != 16 && targetBase != 8 {
					fmt.Printf("\n%s[ERROR] ---->> %s '%d' Is not a valid, known or exitsing base Try(2, 10, 16, 8)%s\n", Red, Yellow, targetBase, Reset)
					continue
				}

				resultStr := strconv.FormatInt(result, targetBase)

				fmt.Printf("%s\n========== CONVERSION RESULT ==========%s\n", Blue, Reset)
				fmt.Printf("  %sHEX VALUE:   %s'%s'%s\n", Blue, Yellow, hexInput, Reset)
				fmt.Printf("  %sBASE RESULT: %s'%d': %s%s\n", Blue, Yellow, targetBase, resultStr, Reset)
				fmt.Printf("%s=================================%s\n", Blue, Reset)

				break
			}
		}

		//("================= [BINARY SECTION]=================")
		if choice == 2 {
			for {
				fmt.Printf("%s\n------ BINARY CONVERSION ------%s\n", Red, Reset)

				fmt.Printf("%sEnter Binary value: %s", Yellow, Reset)
				binInput, _ := reader.ReadString('\n')
				binInput = strings.TrimSpace(binInput)

				result, err := strconv.ParseInt(binInput, 2, 64)
				if err != nil {
					fmt.Printf("%s\n[ERROR] ---> %sInvalid binary number!%s\n", Red, Yellow, Reset)
					continue
				}

				fmt.Printf("%sEnter base value: %s", Yellow, Reset)
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)

				convertTo, err := strconv.Atoi(input)
				if err != nil {
					fmt.Printf("%s\n[INPUT ERROR] ---->> %sEnter a valid base number%s\n", Red, Yellow, Reset)
					continue
				}

				if convertTo != 10 && convertTo != 16 && convertTo != 2 && convertTo != 8 {
					fmt.Printf("%s\n[ERROR] --->> %s '%d' is not a known, valid or existing base, Try(10, 16, 2, 8)%s\n", Red, Yellow, convertTo, Reset)
					continue
				}

				resultStr := strconv.FormatInt(result, convertTo)

				fmt.Printf("%s\n========== RESULT ==========%s\n", Yellow, Reset)
				fmt.Printf("  %sBINARY VALUE:   %s'%s'%s\n", Blue, Yellow, binInput, Reset)
				fmt.Printf("  %sBASE RESULT: %s'%d': %s\n", Blue, Yellow, convertTo, resultStr)
				fmt.Print("============================\n")

				break

			}
		}

		//("=================  [DECIMAL SECTION] =================")
		if choice == 3 {
			for {
				fmt.Printf("%s\n------ DECIMAL CONVERSION ------%s\n", Blue, Reset)

				fmt.Printf("%sEnter Decimal value: %s", Yellow, Reset)
				decInput, _ := reader.ReadString('\n')
				decInput = strings.TrimSpace(decInput)

				result, err := strconv.ParseInt(decInput, 10, 64)
				if err != nil {
					fmt.Printf("\n%s[INPUT ERROR] ----> %s '%s' is not a valid decimal code or value!%s\n", Red, Yellow, decInput, Reset)
					continue
				}

				fmt.Printf("%sEnter the base number: %s", Yellow, Reset)
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)

				convertTo, err := strconv.Atoi(input)
				if err != nil {
					fmt.Printf("\n%s[ERROR] ---->> %sEnter a valid base number%s\n", Red, Yellow, Reset)
					continue
				}

				if convertTo != 2 && convertTo != 16 && convertTo != 10 && convertTo != 8 {
					fmt.Printf("%s\n[INPUT ERROR] ---->> %s '%d' is not a known, valid or exitsing base: Try(2, 16, 10, 8)%s\n", Red, Yellow, convertTo, Reset)
					continue
				}

				resultStr := strconv.FormatInt(result, convertTo)

				fmt.Printf("%s\n========== RESULT ==========%s\n", Blue, Reset)
				fmt.Printf("  %sDECIMAL VALUE:  %s%s%s\n", Blue, Yellow, decInput, Reset)
				fmt.Printf("  %sBASE RESULT: %s'%d': %s%s\n", Blue, Yellow, convertTo, resultStr, Reset)
				fmt.Printf("%s===========================%s\n", Blue, Reset)

				break
			}
		}
	}
}
