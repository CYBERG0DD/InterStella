package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	text     string
	completed bool
}

func displayMenu() {
	fmt.Println("\n--- Simple To-Do List ---")
	fmt.Println("1. View all tasks")
	fmt.Println("2. Add a new task")
	fmt.Println("3. Mark a task as completed")
	fmt.Println("4. Remove a task")
	fmt.Println("5. Exit")
	fmt.Println("***********************\n")
}

func viewTasks(tasks []Task) {
	fmt.Println("\n--- Your Tasks ---")
	if len(tasks) == 0 {
		fmt.Println("You have no tasks in your list.")
	} else {
		for i, task := range tasks {
			status := " "
			if task.completed {
				status = "x"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, status, task.text)
		}
	}
	fmt.Println("******************\n")
}

func addTask(tasks *[]Task) {
	fmt.Print("Enter the new task: ")
	reader := bufio.NewReader(os.Stdin)
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
	
	if task != "" {
		*tasks = append(*tasks, Task{text: task, completed: false})
		fmt.Printf("Task '%s' added successfully.\n", task)
	} else {
		fmt.Println("Task cannot be empty. Please try again.")
	}
}

func markCompleted(tasks []Task) {
	viewTasks(tasks)
	fmt.Print("Enter the number of the task to mark as completed: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}
	
	if num >= 1 && num <= len(tasks) {
		index := num - 1
		tasks[index].completed = true
		fmt.Printf("Task '%s' marked as completed.\n", tasks[index].text)
	} else {
		fmt.Println("Invalid task number. Please try again.")
	}
}

func removeTask(tasks *[]Task) {
	viewTasks(*tasks)
	fmt.Print("Enter the number of the task to remove: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}
	
	if num >= 1 && num <= len(*tasks) {
		removed := (*tasks)[num-1].text
		*tasks = append((*tasks)[:num-1], (*tasks)[num:]...)
		fmt.Printf("Task '%s' removed successfully.\n", removed)
	} else {
		fmt.Println("Invalid task number. Please try again.")
	}
}

func main() {
	var tasks []Task
	
	for {
		displayMenu()
		fmt.Print("Enter your choice (1-5): ")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		
		switch choice {
		case "1":
			viewTasks(tasks)
		case "2":
			addTask(&tasks)
		case "3":
			markCompleted(tasks)
		case "4":
			removeTask(&tasks)
		case "5":
			fmt.Println("Exiting application. Goodbye!!!")
			return
		default:
			fmt.Println("Invalid. Please enter a number.")
		}
	}
}
