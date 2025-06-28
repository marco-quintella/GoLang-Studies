package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
			taskList.ShowProgress()
		case "6":
			taskList.ShowDetailedProgress()
		case "7":
			setDueDateFlow(taskList, scanner)
		case "8":
			searchTasksFlow(taskList, scanner)
		case "9":
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
	fmt.Println("4. 🗑️ Remove task")
	fmt.Println("5. 📊 Show progress")
	fmt.Println("6. 📈 Show detailed progress")
	fmt.Println("7. 📅 Set due date")
	fmt.Println("8. � Search tasks")
	fmt.Println("9. �💾 Save and exit")
}

func addTaskFlow(taskList *TaskList, scanner *bufio.Scanner) {
	fmt.Print("Type the task description: ")
	if scanner.Scan() {
		description := strings.TrimSpace(scanner.Text())
		if description == "" {
			fmt.Println("❌ Description cannot be empty!")
			return
		}

		fmt.Print("Type the task category: ")
		if scanner.Scan() {
			category := strings.TrimSpace(scanner.Text())
			if category == "" {
				fmt.Println("❌ Category cannot be empty!")
				return
			}

			fmt.Print("Type the task priority: ")
			if scanner.Scan() {
				priority := strings.TrimSpace(scanner.Text())
				if priority == "" {
					fmt.Println("❌ Priority cannot be empty!")
					return
				}

				taskList.AddTask(description, category, priority)
				taskList.SaveToFile()
			}
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

func setDueDateFlow(taskList *TaskList, scanner *bufio.Scanner) {
	fmt.Print("Type the ID of the task to set the due date: ")
	var id int
	var err error
	if scanner.Scan() {
		idStr := strings.TrimSpace(scanner.Text())
		id, err = strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("❌ Invalid ID! Type a number.")
			return
		}
	}

	fmt.Print("Type the due date (YYYY-MM-DD): ")
	var dueDate time.Time
	if scanner.Scan() {
		dueDateStr := strings.TrimSpace(scanner.Text())
		dueDate, err = time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			fmt.Println("❌ Invalid date format! Use YYYY-MM-DD.")
			return
		}
	}

	err = taskList.SetDueDate(id, dueDate)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Due date set successfully!")
		taskList.SaveToFile()
	}
}

func searchTasksFlow(taskList *TaskList, scanner *bufio.Scanner) {
	fmt.Println("🔍 Search tasks:")
	fmt.Println("1. By name")
	fmt.Println("2. By status")
	fmt.Print("Choose an option: ")
	if scanner.Scan() {
		option := strings.TrimSpace(scanner.Text())
		if option == "1" {
			searchByNameFlow(taskList, scanner)
		} else if option == "2" {
			searchByStatusFlow(taskList, scanner)
		} else {
			fmt.Println("❌ Invalid option! Type 1 or 2.")
			return
		}
	}
}

func searchByNameFlow(taskList *TaskList, scanner *bufio.Scanner) {
	fmt.Print("Type the search query: ")
	if scanner.Scan() {
		query := strings.TrimSpace(scanner.Text())
		results := taskList.SearchTasks(query)
		if len(results) == 0 {
			fmt.Println("🔍 No tasks found")
			return
		}

		fmt.Println("🔍 Search results:")
		fmt.Println("===================")
		for _, task := range results {
			status := "❌ Pending"
			if task.Completed {
				status = "✅ Completed"
			}
			Colors.Pending.Printf("[%d] %s - %s\n", task.ID, task.Description, status)
			fmt.Printf("    Category: %s\n", task.Category)
		}
		fmt.Println("===================")
	}
}

func searchByStatusFlow(taskList *TaskList, scanner *bufio.Scanner) {
	fmt.Println("🔍 Search by status:")
	fmt.Println("1. ✅ Completed")
	fmt.Println("2. ❌ Pending")
	fmt.Print("Choose an option: ")
	if scanner.Scan() {
		option := strings.TrimSpace(scanner.Text())

		var results []Task
		if option == "1" {
			results = taskList.SearchByStatus(true)
			fmt.Println("Completed tasks:")
		} else if option == "2" {
			results = taskList.SearchByStatus(false)
			fmt.Println("Pending tasks:")
		} else {
			fmt.Println("❌ Invalid option! Type 1 or 2.")
			return
		}

		if len(results) == 0 {
			fmt.Println("🔍 No tasks found")
			return
		}

		for _, task := range results {
			status := "❌ Pending"
			if task.Completed {
				status = "✅ Completed"
			}
			Colors.Pending.Printf("[%d] %s - %s\n", task.ID, task.Description, status)
		}
	}
}
