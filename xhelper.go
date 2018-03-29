package main

import "github.com/fatih/color"

// Let'em be colors
var red = color.New(color.FgRed).PrintfFunc()
var info = color.New(color.Bold, color.FgBlue).PrintlnFunc()
var success = color.New(color.Bold, color.FgGreen).PrintlnFunc()


