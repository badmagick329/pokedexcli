package clicommands

import (
	"fmt"
	"os"
	"time"

	"github.com/badmagick329/pokedexcli/pokeapi"
)

var client = pokeapi.NewClient(time.Hour)

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
		return err
	}
	for _, area := range resp.Results {
		fmt.Printf("- %v: %v\n", area.Name, area.URL)
	}
	return nil
}

func commandMapb() error {
	resp, err := client.ListLocationAreas(true)
	if err != nil {
		return err
	}
	for _, area := range resp.Results {
		fmt.Printf("- %v: %v\n", area.Name, area.URL)
	}
	return nil
}
