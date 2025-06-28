package main

import (
	"fmt"
	"time"
)

type Command interface {
	Execute() error
	Undo() error
}

type AddTaskCommand struct {
	taskList    *TaskList
	description string
	category    string
	priority    Priority
	taskID      int
	wasExecuted bool
}

func NewAddTaskCommand(taskList *TaskList, description string, category string, priority Priority) *AddTaskCommand {
	return &AddTaskCommand{
		taskList:    taskList,
		description: description,
		category:    category,
		priority:    priority,
		wasExecuted: false,
	}
}

func (c *AddTaskCommand) Execute() error {
	if !c.wasExecuted {
		// First time executing - create new task
		c.taskID = c.taskList.NextID
		c.taskList.AddTask(c.description, c.category, c.priority.String())
		c.wasExecuted = true
	} else {
		// Redo - re-add the task with the same ID
		task := Task{
			ID:          c.taskID,
			Description: c.description,
			Category:    c.category,
			Priority:    c.priority.String(),
			Completed:   false,
			CreatedAt:   time.Now(),
		}
		c.taskList.Tasks = append(c.taskList.Tasks, task)

		// Update NextID if needed
		if c.taskID >= c.taskList.NextID {
			c.taskList.NextID = c.taskID + 1
		}

		fmt.Printf("✅ Task re-added: %s (ID: %d)\n", c.description, c.taskID)
	}
	return nil
}

func (c *AddTaskCommand) Undo() error {
	return c.taskList.RemoveTask(c.taskID)
}

type CompleteTaskCommand struct {
	taskList            *TaskList
	taskID              int
	wasCompleted        bool
	originalCompletedAt *time.Time
}

func NewCompleteTaskCommand(taskList *TaskList, taskID int) *CompleteTaskCommand {
	return &CompleteTaskCommand{
		taskList: taskList,
		taskID:   taskID,
	}
}

func (c *CompleteTaskCommand) Execute() error {
	for i := range c.taskList.Tasks {
		if c.taskList.Tasks[i].ID == c.taskID {
			c.wasCompleted = c.taskList.Tasks[i].Completed
			c.originalCompletedAt = c.taskList.Tasks[i].CompletedAt
			break
		}
	}
	return c.taskList.CompleteTask(c.taskID)
}

func (c *CompleteTaskCommand) Undo() error {
	if c.wasCompleted {
		// Task was already completed, restore original state
		for i := range c.taskList.Tasks {
			if c.taskList.Tasks[i].ID == c.taskID {
				c.taskList.Tasks[i].Completed = true
				c.taskList.Tasks[i].CompletedAt = c.originalCompletedAt
				fmt.Printf("✅ Task %d restored to completed state\n", c.taskID)
				return nil
			}
		}
	} else {
		// Task was not completed, mark as incomplete
		for i := range c.taskList.Tasks {
			if c.taskList.Tasks[i].ID == c.taskID {
				c.taskList.Tasks[i].Completed = false
				c.taskList.Tasks[i].CompletedAt = nil
				fmt.Printf("↩️ Task %d marked as incomplete\n", c.taskID)
				return nil
			}
		}
	}

	return fmt.Errorf("task %d not found", c.taskID)
}

type CommandManager struct {
	history   []Command
	redoStack []Command
}

func NewCommandManager() *CommandManager {
	return &CommandManager{
		history:   make([]Command, 0),
		redoStack: make([]Command, 0),
	}
}

func (cm *CommandManager) Execute(cmd Command) error {
	err := cmd.Execute()
	if err != nil {
		return err
	}
	cm.history = append(cm.history, cmd)
	// Clear redo stack when new command is executed
	cm.redoStack = make([]Command, 0)
	return nil
}

func (cm *CommandManager) Undo() error {
	if len(cm.history) == 0 {
		return fmt.Errorf("no commands to undo")
	}

	lastIndex := len(cm.history) - 1
	lastCmd := cm.history[lastIndex]

	err := lastCmd.Undo()
	if err != nil {
		return err
	}

	cm.redoStack = append(cm.redoStack, lastCmd)
	cm.history = cm.history[:lastIndex]
	return nil
}

func (cm *CommandManager) Redo() error {
	if len(cm.redoStack) == 0 {
		return fmt.Errorf("no commands to redo")
	}

	lastIndex := len(cm.redoStack) - 1
	lastCmd := cm.redoStack[lastIndex]

	err := lastCmd.Execute()
	if err != nil {
		return err
	}

	cm.history = append(cm.history, lastCmd)
	cm.redoStack = cm.redoStack[:lastIndex]
	return nil
}
