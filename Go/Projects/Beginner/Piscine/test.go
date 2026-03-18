package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Standardize to capital 'T' if you want to use Task elsewhere
type Task struct {
	ID   int
	Note string
	Done bool
}

func main() {
	var todoList []Task
	reader := bufio.NewReader(os.Stdin)
	count := 1

	for {
		fmt.Println("\n--- GO TO DO LIST ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks") // Fixed numbering
		fmt.Println("3. Exit")

		fmt.Print("Select choice: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input", err)
			continue
		}
		choice := strings.TrimSpace(input)

		if choice == "" {
			fmt.Println("Please fill here...")
			continue
		}

		// Keep choice as a string to make the switch cases easier
		switch choice {
		case "1":
			fmt.Print("Add new task: ")
			note, _ := reader.ReadString('\n')
			note = strings.TrimSpace(note)

			if note != "" {
				newTask := Task{ID: count, Note: note, Done: false}
				todoList = append(todoList, newTask)
				count++
				fmt.Println("New Task added 👍!")
			}
		case "2":
			fmt.Println("\nYOUR TASKS:")
			if len(todoList) == 0 {
				fmt.Println("[List is currently empty]")
			}
			for _, t := range todoList {
				status := "[ ]"
				if t.Done {
					status = "[X]"
				}
				fmt.Printf("%d. %s %s\n", t.ID, status, t.Note)
			}
		case "3":
			fmt.Println("Goodbye!")
			return // This exits the whole program

		default:
			fmt.Println("Invalid choice, try again")
		}
	}
}