package main

import "strings"

func (tl *TaskList) SearchTasks(query string) []Task {
	if query == "" {
		return tl.Tasks
	}

	query = strings.ToLower(query)
	var results []Task

	for _, task := range tl.Tasks {
		if strings.Contains(strings.ToLower(task.Description), query) {
			results = append(results, task)
		}
	}

	return results
}

func (tl *TaskList) SearchByStatus(completed bool) []Task {
	var results []Task

	for _, task := range tl.Tasks {
		if task.Completed == completed {
			results = append(results, task)
		}
	}

	return results
}
