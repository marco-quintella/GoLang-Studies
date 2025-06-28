package main

import "github.com/fatih/color"

type Color struct {
	Title    *color.Color
	Complete *color.Color
	Pending  *color.Color
	Warning  *color.Color
	Error    *color.Color
}

// Save colors in a global variable
var Colors = &Color{
	Title:    color.New(color.FgCyan, color.Bold),
	Complete: color.New(color.FgGreen, color.Bold),
	Pending:  color.New(color.FgYellow, color.Bold),
	Warning:  color.New(color.FgYellow, color.Bold),
	Error:    color.New(color.FgRed, color.Bold),
}
