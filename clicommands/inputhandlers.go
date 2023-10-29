package clicommands

import (
	"fmt"
	"strings"
)

func HandleInput(inp string, commands map[string]CliCommand) string {
	cleaned := cleanInput(inp)
	cmd, ok := commands[strings.Join(cleaned, "")]
	if !ok {
		return fmt.Sprintf("%v is not a valid command. Type 'help' to display command list\n", inp)
	}
	err := cmd.callback()
	if err != nil {
		return fmt.Sprintln("Encountered error: ", err)
	}
	return ""
}

func cleanInput(inp string) []string {
	return strings.Fields(strings.ToLower(inp))
}
