package commands

// GetCommands returns a dictionary of available commands.
func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists next location areas",
			Callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous location areas",
			Callback:    CommandMapb,
		},
		"explore": {
			name:        "explore",
			params:      []string{"area"},
			description: "Lists pokemon in a certain area",
			Callback:    CommandExplore,
		},
	}
}
