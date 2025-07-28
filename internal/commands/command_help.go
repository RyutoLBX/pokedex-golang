package commands

import (
	"fmt"
)

// Prints out a list of commands and their descriptions.
func CommandHelp(_ *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	defer fmt.Println()

	commands := GetCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
