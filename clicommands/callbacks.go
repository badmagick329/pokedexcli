package clicommands

import (
	"fmt"
	"log"
	"os"

	"github.com/badmagick329/pokedexcli/pokeapi"
)

var client = pokeapi.NewClient()

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

func commandMap() error {
	resp, err := client.ListLocationAreas(false)
	if err != nil {
		log.Fatal(nil)
	}
	for _, area := range resp.Results {
		fmt.Printf("- %v: %v\n", area.Name, area.URL)
	}
	return nil
}

func commandMapb() error {
	resp, err := client.ListLocationAreas(true)
	if err != nil {
		log.Fatal(nil)
	}
	for _, area := range resp.Results {
		fmt.Printf("- %v: %v\n", area.Name, area.URL)
	}
	return nil
}
