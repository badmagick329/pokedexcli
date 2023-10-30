package clicommands

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/badmagick329/pokedexcli/pokeapi"
)

var config = pokeapi.Config{
	Pokemons: make(map[string]pokeapi.Pokemon),
}
var client = pokeapi.NewClient(time.Hour, config)

func Help() error {
	fmt.Println("Usage:")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func Exit() error {
	fmt.Println("Goodbye 👋")
	os.Exit(0)
	return nil
}

func Map() error {
	resp, err := client.ListLocationAreas(false)
	if err != nil {
		return err
	}
	for _, area := range resp.Results {
		fmt.Printf("- %v: %v\n", area.Name, area.URL)
	}
	return nil
}

func Mapb() error {
	resp, err := client.ListLocationAreas(true)
	if err != nil {
		return err
	}
	for _, area := range resp.Results {
		fmt.Printf("- %v: %v\n", area.Name, area.URL)
	}
	return nil
}

func Explore(s []string) error {
	if len(s) == 0 {
		return fmt.Errorf("No argument received for exploration")
	}
	resp, err := client.LocationDetails(s[0])
	if err != nil {
		return err
	}
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf("Pokemon: %v\n", encounter.Pokemon.Name)
	}
	return nil
}

func Catch(s []string) error {
	if len(s) == 0 {
		return fmt.Errorf("No argument received for catch")
	}
	name := s[0]
	resp, err := client.CatchPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	const threshold = 50
	randomNum := rand.Intn(resp.BaseExperience)
	if randomNum > threshold {
		fmt.Printf("Failed to catch %s\n", name)
		return nil
	}
	fmt.Printf("%s caught!\n", name)
	config.Pokemons[name] = resp
	return nil
}

func Inspect(s []string) error {
	if len(s) == 0 {
		return fmt.Errorf("No argument received for inspect")
	}
	name := s[0]
	pokemon, ok := config.Pokemons[name]
	if !ok {
		fmt.Println("You have no caught this pokemon")
		return nil
	}
	printPokemon(name, pokemon)
	return nil
}

func printPokemon(name string, pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Experience: %v\n", pokemon.BaseExperience)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  -%s\n", pokemonType.Type.Name)
	}

}

func Pokedex(s []string) error {
	if len(config.Pokemons) == 0 {
		fmt.Println("You have not caught any pokemons")
		return nil
	}
	fmt.Println("Your pokedex: ")
	for _, pokemon := range config.Pokemons {
		fmt.Printf("  -%s\n", pokemon.Name)
	}
	return nil
}
