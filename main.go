package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/badmagick329/pokedexcli/clicommands"
	"github.com/badmagick329/pokedexcli/pokeapi"
)

func main() {
	// run()
	apiTest()
}

func apiTest() {
	url := "https://pokeapi.co/api/v2/location-area/"
	res := pokeapi.Get(url)
	var data pokeapi.LocationArea
	pokeapi.Json(res, &data)
	fmt.Println(data)
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
