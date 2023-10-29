package main

import (
	"bufio"
	"fmt"
	clicommands "github.com/badmagick329/pokedexcli/clicommands"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			fmt.Printf("pokedex > ")
			continue
		}
		output := clicommands.HandleInput(text)
		if output != "" {
			fmt.Printf(output)
		}
		fmt.Printf("pokedex > ")
	}
}
