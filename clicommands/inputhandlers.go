package clicommands

import (
	"fmt"
)

func HandleInput(inp string) {
	commands := getCommands()
	for key := range commands {
		if key == inp {
			fmt.Printf("%v is a valid command\n", inp)
			return
		}
	}
	fmt.Printf("%v is not a valid command\n", inp)
}
