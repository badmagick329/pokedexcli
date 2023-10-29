package clicommands

import (
	"fmt"
	"os"
)


func commandHelp() error {
	fmt.Println("Usage:")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye 👋")
	os.Exit(0)
	return nil
}

