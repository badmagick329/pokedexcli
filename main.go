package main

import (
	"bufio"
	"fmt"
	clicommands "github.com/badmagick329/pokedexcli/clicommands"
	"os"
)

func main() {
	run()
}

func run() {
	prompt := prompter()
	for {
		text := prompt()
		if text == "" {
			continue
		}
		output := clicommands.HandleInput(text)
		if output != "" {
			fmt.Printf(output)
		}
	}
}

func prompter() func() string {
	scanner := bufio.NewScanner(os.Stdin)
	return func() string {
		fmt.Printf("Pokedex > ")
		res := scanner.Scan()
		if res {
			return scanner.Text()
		} else {
			fmt.Println("An error occured")
			fmt.Println(scanner.Err().Error())
			os.Exit(1)
			return ""
		}
	}
}
