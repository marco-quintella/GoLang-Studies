package main

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

func (tl *TaskList) ShowProgress() {
	total := len(tl.Tasks)
	if total == 0 {
		Colors.Warning.Println("No tasks found")
		return
	}

	completed := 0
	for _, task := range tl.Tasks {
		if task.Completed {
			completed++
		}
	}

	percent := float64(completed) / float64(total) * 100

	// Display basic progress
	fmt.Printf("\n📊 Progress:")
	fmt.Printf("===================\n")
	fmt.Printf("Total tasks: %d\n", total)
	Colors.Complete.Printf("Completed: %d\n", completed)
	Colors.Pending.Printf("Pending: %d\n", total-completed)
	fmt.Printf("Progress: %.2f%%\n", percent)
	fmt.Printf("===================\n")

	bar := progressbar.NewOptions(total,
		progressbar.OptionSetDescription("Task progress..."),
		progressbar.OptionSetWidth(50),
		progressbar.OptionShowCount(),
	)

	bar.Set(completed)
	fmt.Println("")
}

func (tl *TaskList) ShowDetailedProgress() {
	total := len(tl.Tasks)
	if total == 0 {
		Colors.Warning.Println("No tasks found")
		return
	}

	completed := 0
	pending := 0

	for _, task := range tl.Tasks {
		if task.Completed {
			completed++
		} else {
			pending++
		}
	}

	fmt.Printf("\n📈 Detailed Progress:\n")
	fmt.Printf("===================\n")

	if completed > 0 {
		fmt.Printf("Completed: %d\n", completed)
		completedBar := progressbar.NewOptions(total,
			progressbar.OptionSetDescription("Completed tasks..."),
			progressbar.OptionSetWidth(40),
			progressbar.OptionShowCount(),
		)
		completedBar.Set(completed)
		fmt.Println()
	}

	if pending > 0 {
		fmt.Printf("Pending: %d\n", pending)
		pendingBar := progressbar.NewOptions(total,
			progressbar.OptionSetDescription("Pending tasks..."),
			progressbar.OptionSetWidth(40),
			progressbar.OptionShowCount(),
		)
		pendingBar.Set(pending)
		fmt.Println()
	}

	fmt.Println()
}
