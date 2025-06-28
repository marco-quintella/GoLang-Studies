package main

import "strings"

type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
)

func StringToPriority(s string) Priority {
	switch strings.ToLower(s) {
	case "low":
		return Low
	case "medium":
		return Medium
	case "high":
		return High
	default:
		return Low
	}
}

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	default:
		return "Unknown"
	}
}
