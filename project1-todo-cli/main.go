package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Load existing tasks
	taskList, err := LoadFromFile()
	if err != nil {
		// If the file does not exist, create a new task list
		fmt.Printf("Error loading tasks: %v\n", err)
		taskList = NewTaskList()
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🎯 TODO CLI - Task Manager")
	fmt.Println("====================================")

	for {
		showMenu()
		fmt.Print("Choose an option: ")

		if !scanner.Scan() {
			break
		}

		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			addTaskFlow(taskList, scanner)
		case "2":
			taskList.ListTasks()
		case "3":
			completeTaskFlow(taskList, scanner)
		case "4":
			removeTaskFlow(taskList, scanner)
		case "5":
			fmt.Println("💾 Saving tasks...")
			err := taskList.SaveToFile()
			if err != nil {
				fmt.Printf("❌ Error saving tasks: %v\n", err)
			} else {
				fmt.Println("💾 Tasks saved successfully!")
			}
			return
		default:
			fmt.Println("❌ Invalid option! Type a number between 1 and 5.")
		}

		fmt.Println() // Add a blank line for better readability
	}
}

func showMenu() {
	fmt.Println("\n📋 MAIN MENU:")
	fmt.Println("1. ➕ Add task")
	fmt.Println("2. 📝 List tasks")
	fmt.Println("3. ✅ Mark as completed")
	fmt.Println("4. 🗑️  Remove task")
	fmt.Println("5. 💾 Save and exit")
}

func addTaskFlow(taskList *TaskList, scanner *bufio.Scanner) {
	fmt.Print("Type the task description: ")
	if scanner.Scan() {
		description := strings.TrimSpace(scanner.Text())
		if description != "" {
			taskList.AddTask(description)
			// Save tasks to file
			taskList.SaveToFile()
		} else {
			fmt.Println("❌ Description cannot be empty!")
		}
	}
}

func completeTaskFlow(taskList *TaskList, scanner *bufio.Scanner) {
	fmt.Print("Type the ID of the task to mark as completed: ")
	if scanner.Scan() {
		idStr := strings.TrimSpace(scanner.Text())
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("❌ Invalid ID! Type a number.")
			return
		}

		err = taskList.CompleteTask(id)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		} else {
			taskList.SaveToFile()
		}
	}
}

func removeTaskFlow(taskList *TaskList, scanner *bufio.Scanner) {
	fmt.Print("Type the ID of the task to remove: ")
	if scanner.Scan() {
		idStr := strings.TrimSpace(scanner.Text())
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("❌ Invalid ID! Type a number.")
			return
		}

		err = taskList.RemoveTask(id)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		} else {
			taskList.SaveToFile()
		}
	}
}
