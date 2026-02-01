package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("WELCOME TO GLOBAL \"9\" CORPORATION.")
	fmt.Println("[Motto: With GOD - Beyond Limits]")
	fmt.Println("\nPlease fill in the required information below so we can begin processing your details.\n")

	reader := bufio.NewReader(os.Stdin)

	// --- BASIC INFO ---
	fmt.Print("Enter first name: ")
	firstName := readInput(reader)
	fmt.Print("Middle name [optional]: ")
	middleName := readInput(reader)
	fmt.Print("Enter last name: ")
	lastName := readInput(reader)

	fmt.Printf("\nHello, %s %s %s, and welcome aboard!\n\n", strings.Title(firstName), strings.Title(middleName), strings.Title(lastName))

	// --- ACCOUNT CREATION ---
	for {
		fmt.Print("Enter username or [nickname]: ")
		username := readInput(reader)
		if len(username) < 5 {
			fmt.Println("Invalid. Username must be at least 5 characters long.")
			continue
		}

		fmt.Print("Enter email address: ")
		email := readInput(reader)
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
			fmt.Println("Invalid email address. Please try again.")
			continue
		}

		// --- PHONE ---
		var phone string
		for {
			fmt.Print("Enter phone number (11 digits): ")
			phone = readInput(reader)
			if len(phone) != 11 {
				fmt.Println("Invalid. Must be exactly 11 digits.")
			} else if !isDigits(phone) {
				fmt.Println("Invalid. Phone must contain only numbers.")
			} else {
				break
			}
		}

		// --- PASSWORD ---
		fmt.Print("Suggest strong password? (yes/no): ")
		suggest := strings.ToLower(readInput(reader))

		if suggest == "yes" {
			fmt.Println("Suggested password: I7baS3%&*")
			fmt.Println("You can save your password to your device or Google for recovery.")
		} else {
			fmt.Println("We recommend using a strong password — your account could be vulnerable otherwise.")
			for {
				fmt.Print("Create password (8–10 characters): ")
				password := readInput(reader)
				if len(password) < 8 {
					fmt.Println("Invalid. Password should be at least 8 characters.")
				} else if len(password) > 10 {
					fmt.Println("Invalid. Password should not exceed 10 characters.")
				} else {
					fmt.Println("Password accepted.")
					break
				}
			}
		}

		// --- ID VALIDATION ---
		fmt.Println("\nYour ID will be sent to your email. Don't share it with anyone.")
		fmt.Println("Note: It cannot be changed once issued.\n")

		// --- AGE VALIDATION ---
		var age int
		for {
			fmt.Print("Enter birth year (YYYY): ")
			birthYear := readInput(reader)
			year, err := strconv.Atoi(birthYear)
			if err != nil {
				fmt.Println("Please enter a valid year.")
				continue
			}
			age = 2025 - year
			fmt.Printf("You are %d years old.\n", age)
			if age < 18 || age > 45 {
				fmt.Println("Sorry, your age does not meet our requirement (18–45).")
				continue
			}
			break
		}

		// --- GENDER ---
		fmt.Print("Gender (M/F/Other): ")
		gender := strings.Title(readInput(reader))

		// --- PROFESSION ---
		professions := []string{
			"cyber security specialist",
			"software engineer",
			"hardware engineer",
			"it technician",
			"financial management",
			"customer care",
		}

		fmt.Println("\nFIELD OF OPERATION:")
		for _, p := range professions {
			fmt.Printf("- %s\n", strings.Title(p))
		}

		var choice string
		for {
			fmt.Print("\nEnter your profession exactly as listed: ")
			choice = strings.ToLower(readInput(reader))
			if !contains(professions, choice) {
				fmt.Println("Sorry, that profession is not on the list. Please check your spelling.")
				continue
			}
			break
		}

		// --- STRENGTH MARK ---
		var mark int
		for {
			fmt.Print("Rate your field strength (100, 95, 90, 85, 80, 75, 70): ")
			markStr := readInput(reader)
			val, err := strconv.Atoi(markStr)
			if err != nil {
				fmt.Println("Please enter a valid number.")
				continue
			}
			if val < 70 {
				fmt.Println("Sorry, that's below our acceptable range.")
				continue
			}
			mark = val
			break
		}

		switch {
		case mark == 100:
			fmt.Println("You are an expert in this field — we need the best of your knowledge to help us grow.")
		case mark >= 95:
			fmt.Println("Excellent knowledge! We trust you'll help in our company's growth.")
		case mark >= 90:
			fmt.Println("Strong skills! You'll be a great asset to our company.")
		case mark >= 85:
			fmt.Println("Good knowledge — your expertise will help us move forward.")
		case mark >= 75:
			fmt.Println("Solid understanding. You'll make our company outstanding.")
		default:
			fmt.Println("Good — we're happy to have you onboard.")
		}

		// --- STUDENT CHECK ---
		fmt.Print("\nAre you a student (yes/no)?: ")
		student := strings.ToLower(readInput(reader))
		if student == "yes" {
			fmt.Println("\nWelcome, young star! You've come to the right place.")
			fmt.Println("We will begin processing your details and contact you shortly.")
			fmt.Println("You will receive an email with your job qualification status.")
		} else {
			fmt.Println("\nThank you for your cooperation. We will begin processing your details shortly.")
		}

		fmt.Println("\n✅ Registration complete! Welcome to GLOBAL '9' CORPORATION.")
		break
	}
}

// --- FIXED HELPER FUNCTIONS ---

// readInput reads a line safely, handles errors
func readInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(input)
}

// isDigits checks if string contains only digits
func isDigits(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

// contains checks case-insensitive match in profession list
func contains(list []string, item string) bool {
	for _, v := range list {
		if strings.ToLower(v) == item {
			return true
		}
	}
	return false
}
