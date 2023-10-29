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
			break
		}
		clicommands.HandleInput(text)
		fmt.Printf("pokedex > ")
	}
}
