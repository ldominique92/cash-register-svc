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

// TODO: add could be better
// add total to cart list
// implement remove
// we don't need a getter for total in the app I guess
