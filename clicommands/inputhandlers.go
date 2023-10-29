package clicommands

import (
	"fmt"
)

func HandleInput(inp string) string {
	commands := getCommands()
	cmd, ok := commands[inp]
	if !ok {
		return fmt.Sprintf("%v is not a valid command. Type 'help' to display command list\n", inp)
	}
	err := cmd.callback()
	if err != nil {
		return fmt.Sprintln("Encountered error: ", err)
	}
	return ""
}
