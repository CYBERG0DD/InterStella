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
		var num1, num2 float64
		var err error

		fmt.Println(Green + "\n=============== Calculator CLI based ==============" + Reset)

		operations := []string{"Addition", "Subtraction", "Multiplication", "Division", "Help", "Exit"}
		for i, op := range operations {
			fmt.Printf("%s[%d]. %s%s\n", Blue, i+1, op, Reset)
		}

		fmt.Printf(Yellow + "Enter your choice (1-6): " + Reset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err = strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 6 {
			fmt.Printf("%s[INVALID]: --->> Please enter a number between %s(1 and 6).%s", Red, Yellow, Reset)
			continue
		}

		if choice == 5 {
			fmt.Print("Need help? Try this\n")
			help := []string{"add <a> + <b>   → addition", "sub <a> - <b>   →  Subtraction", "Multiply <a> * <b>   → multplication", "Divivsion <a> / <b> → Division"}
			for p, d := range help {
				fmt.Printf("%s[%d]. %s%s\n", Blue, p+1, d, Reset)
			}
			continue
		}
		if choice == 6 {
			fmt.Printf("%sNice doing some math%s\n", Red, Reset)
			fmt.Printf("%sEXITING.....Closing Application%s\n", Red, Reset)
			break
		}

		fmt.Print("Enter first number: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num1, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("%s[Invalid] --->> Cannot use string as value in place of int%s\n", Red, Reset)
			continue
		}

		fmt.Print("Enter second number: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num2, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("%s[Invalid] -->> Cannot use string a as value in place of int%s\n", Red, Reset)
			continue
		}

		var result float64

		if choice == 1 {
			result = num1 + num2
			fmt.Printf("You added %s%v%s and %s%v%s: Result: (%v + %v) = %v\n", Green, num1, Reset, Green, num2, Reset, num1, num2, result)
		} else if choice == 2 {
			result = num1 - num2
			fmt.Printf("You subtracted %s%v%s and %s%v%s: Result: (%v - %v) = %v\n", Green, num1, Reset, Green, num2, Reset, num1, num2, result)

		} else if choice == 3 {
			result = num1 * num2
			fmt.Printf("You Multiplied %s%v%s and %s%v%s: Result: (%v * %v) = %v\n", Green, num1, Reset, Green, num2, Reset, num1, num2, result)
		} else if choice == 4 {
			if num2 == 0 {
				fmt.Printf("%s[INVALID OPERATION]: Zero division error. Cannot divide a number by %s0 but can divide %s0 by a number.%s\n", Red, Yellow, Yellow, Reset)
				continue
			}
			result = num1 / num2
			fmt.Printf("You Divided %s%v%s and %s%v%s: Result: (%v / %v) = %v\n", Green, num1, Reset, Green, num2, Reset, num1, num2, result)
		}
	}
}
