/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cash-register-svc/internal/cmd"

	cobraprompt "github.com/stromland/cobra-prompt"
)

var simplePrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	AddDefaultExitCommand:    true,
	DisableCompletionCommand: true,
}

func main() {
	simplePrompt.Run()
}
