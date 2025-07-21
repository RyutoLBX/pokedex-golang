package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	c := &config{Next: nil, Previous: nil}
	c.Next = stringToPtr("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		commandName := words[0]

		// Handle command (first word)
		commands := getCommands()
		command, exists := commands[commandName]
		if exists {
			err := command.callback(c)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists next location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous location areas",
			callback:    commandMapb,
		},
	}
}
