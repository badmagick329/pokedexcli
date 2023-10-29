package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/badmagick329/pokedexcli/clicommands"
)

func main() {
	run()
}

func run() {
	prompt := prompter()
	commands := clicommands.GetCommands()
	for {
		text := prompt()
		if text == "" {
			continue
		}
		output := clicommands.HandleInput(text, commands)
		if output != "" {
			fmt.Printf(output)
		}
	}
}

func prompter() func() string {
	scanner := bufio.NewScanner(os.Stdin)
	return func() string {
		fmt.Printf("Pokedex > ")
		ok := scanner.Scan()
		if ok {
			return scanner.Text()
		} else {
			log.Fatal(scanner.Err().Error())
			return ""
		}
	}
}
