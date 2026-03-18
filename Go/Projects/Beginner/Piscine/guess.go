package main

import (
	"fmt"
	"strings"
)

func main() {
	var set_password string = "INFINITE"
	var password = ""
	var limit int = 5

	for stop := 1; stop <= limit; stop++ { // This counts the number of trials left for the user

		fmt.Print("Guess the word: ") //Takes input from the user
		fmt.Scanln(&password)

		password = strings.TrimSpace(password)

		if password == set_password {
			fmt.Println("Access Granted. Smart kid! 😁")
			break
		} else {
			fmt.Println("Invalid Guess😡! Try again ")
		}
		if limit == stop {
			fmt.Println("Sorry mate you failed!😔 You are out of tries.  ")
		}
	}
}
