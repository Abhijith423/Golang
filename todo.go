package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title     string
	Completed bool
}

var tasks []Task

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- TO DO LIST ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		choiceStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error")
			continue
		}

		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		switch choice {
		case 1:
			addTask(reader)
		case 2:
			viewTasks()
		case 3:
			completeTask(reader)
		case 4:
			deleteTask(reader)
		case 5:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	tasks = append(tasks, Task{Title: title})
	fmt.Println("Task added successfully!")
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	for i, task := range tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}
}

func completeTask(reader *bufio.Reader) {
	viewTasks()
	fmt.Print("Enter task number to mark completed: ")

	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)

	if err != nil || num < 1 || num > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks[num-1].Completed = true
	fmt.Println("Task marked as completed!")
}

func deleteTask(reader *bufio.Reader) {
	viewTasks()
	fmt.Print("Enter task number to delete: ")

	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)

	if err != nil || num < 1 || num > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks = append(tasks[:num-1], tasks[num:]...)
	fmt.Println("Task deleted successfully!")
}
