package main

import (
	"fmt"
	"time"
)

// Single task
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

// Task list
type TaskList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks:  make([]Task, 0),
		NextID: 1,
	}
}

func (tl *TaskList) AddTask(description string) {
	task := Task{
		ID:          tl.NextID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
	fmt.Printf("✅ Task added: %s (ID: %d)\n", description, task.ID)
}

func (tl *TaskList) CompleteTask(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			if tl.Tasks[i].Completed {
				return fmt.Errorf("task %d already completed", id)
			}
			tl.Tasks[i].Completed = true
			now := time.Now()
			tl.Tasks[i].CompletedAt = &now
			fmt.Printf("✅ Task %d marked as completed!\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (tl *TaskList) RemoveTask(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			fmt.Printf("🗑️ Task %d removed!\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (tl *TaskList) ListTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("🔍 No tasks found")
		return
	}

	fmt.Println("🔍 Task List:")
	fmt.Println("===================")

	for _, task := range tl.Tasks {
		status := "❌ Pending"
		if task.Completed {
			status = "✅ Completed"
		}

		if task.Completed {
			Colors.Complete.Printf("[%d] %s - %s\n", task.ID, task.Description, status)
			fmt.Printf("    Created at: %s\n", task.CreatedAt.Format("02/01/2006 15:04"))
			fmt.Printf("    Completed at: %s\n", task.CompletedAt.Format("02/01/2006 15:04"))
		} else {
			Colors.Pending.Printf("[%d] %s - %s\n", task.ID, task.Description, status)
			fmt.Printf("    Created at: %s\n", task.CreatedAt.Format("02/01/2006 15:04"))
		}

		fmt.Println("===================")
	}
}
