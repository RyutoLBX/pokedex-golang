package repl

import (
	"bufio"
	"fmt"
	"os"

	cm "github.com/RyutoLBX/pokedexcli/internal/commands"
	hp "github.com/RyutoLBX/pokedexcli/internal/helpers"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	c := &cm.Config{Next: nil, Previous: nil}
	c.Next = hp.Ptr("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := hp.CleanInput(scanner.Text())

		commandName := words[0]

		// Handle command (first word)
		commands := cm.GetCommands()
		command, exists := commands[commandName]
		if exists {
			err := command.Callback(c)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}
