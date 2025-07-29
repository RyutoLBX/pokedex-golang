package commands

import (
	"fmt"
)

// CommandHelp prints out a list of commands and their descriptions.
func CommandHelp(_ *Config, _ []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	defer fmt.Println()

	commands := GetCommands()
	for _, command := range commands {
		fmt.Printf("%s", command.name)
		for _, param := range command.params {
			fmt.Printf(" [%s]", param)
		}
		fmt.Printf(": %s\n", command.description)
	}
	return nil
}
