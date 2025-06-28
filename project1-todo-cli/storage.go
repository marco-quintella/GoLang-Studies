package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename = "tasks.json"

func (tl *TaskList) SaveToFile() error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return fmt.Errorf("error converting to JSON: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}

func LoadFromFile() (*TaskList, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return NewTaskList(), nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var taskList TaskList
	err = json.Unmarshal(data, &taskList)
	if err != nil {
		return nil, fmt.Errorf("error converting JSON to struct: %v", err)
	}

	return &taskList, nil
}
