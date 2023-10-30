package clicommands

import (
	"fmt"
	"log"
	"strings"
)

func HandleInput(inp string, commands map[string]CliCommand) string {
	cleaned := cleanInput(inp)
	if len(cleaned) == 0 {
		log.Fatalln("Empty commands should not be accepted")
	}
	cmd, ok := commands[cleaned[0]]
	if !ok {
		return fmt.Sprintf("%v is not a valid command. Type 'help' to display command list\n", inp)
	}
	err := cmd.callback(cleaned[1:])
	if err != nil {
		return fmt.Sprintln("Encountered error: ", err)
	}
	return ""
}

func cleanInput(inp string) []string {
	return strings.Fields(strings.ToLower(inp))
}
