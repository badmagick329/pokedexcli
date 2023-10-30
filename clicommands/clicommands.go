package clicommands

type CliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func(s []string) error { return Help() },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func(s []string) error { return Exit() },
		},
		"map": {
			name:        "map",
			description: "List next locations",
			callback:    func(s []string) error { return Map() },
		},
		"mapb": {
			name:        "mapb",
			description: "List previous locations",
			callback:    func(s []string) error { return Mapb() },
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore a location",
			callback:    func(s []string) error { return Explore(s) },
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Catch a pokemon",
			callback:    func(s []string) error { return Catch(s) },
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect a pokemon you've caught",
			callback:    func(s []string) error { return Inspect(s) },
		},
	}
}
