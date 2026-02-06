package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type task struct{
		ID int
		Note string
		Done bool
		
	}

func main() {

	var todoLisst []Task
	reader := bufio.NewReader(os.Stdin)
	count := 1
	 
	for {
		fmt.Println("\n --- GO TO DO LIST ---")
		fmt.Println("1. Add Task")
		fmt.Println("1. View Task")
		fmt.Println("1. Exit")

		input, err := reader.Readstring('\n')
		if err != nil{
			fmt.Println("Error reading input", err)
		}
		choice := strings.TrimSpace(input)

		if choice == ""{
			fmt.Println("Please fill here...")
			continue
		}
		
	}
	switch choice{
	case "1":
		fmt.Print("Add new task: ")
		note, _ := reader.ReadString('\n')
		note = sttrings.TrimSpace(note)

		if note != ""{
			newTask := Task{ID: count, Nte: note, Done: false}
			todoList = append(todoList, newTask)
			count++
			fmt.Println("New Task added 👍!")
		}
	case "2":
		fmt.Println("\nYOUR TASKS: ")
		if len(todoList) == 0{
			fmt.Println("[List is currently empty]")
		}
		for _, t := range todoList{
			status := "[ ]"
			if t.Done{
				status = "[X]"
			}
			fmt.Prinf("%d. %s %s\n", t.ID, status, t.Note)
		}
	case "3":
		fmt.Println("Goodbye!")
		return

	default:
		fmt.Println("Invalid choice, try again")
		

	}
}
