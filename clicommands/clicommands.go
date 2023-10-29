package clicommands

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous locations",
			callback:    commandMapb,
		},
	}
}
