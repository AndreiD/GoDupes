package main

import "github.com/fatih/color"

// Let'em be colors

var info = color.New(color.Bold, color.FgBlue).PrintfFunc()
var success = color.New(color.Bold, color.FgGreen).PrintfFunc()
var red = color.New(color.FgRed).PrintfFunc()
var warn = color.New(color.FgHiMagenta).PrintfFunc()

