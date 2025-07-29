package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	cm "github.com/RyutoLBX/pokedexcli/internal/commands"
	hp "github.com/RyutoLBX/pokedexcli/internal/helpers"
	pkc "github.com/RyutoLBX/pokedexcli/internal/pokecache"
)

// StartRepl initiates the PokedexCLI REPL loop.
func StartRepl() {
	// Scanner init
	scanner := bufio.NewScanner(os.Stdin)

	// Config init
	config := &cm.Config{Next: nil, Previous: nil}
	config.Next = hp.Ptr("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")

	// Map cache init
	mapCache := pkc.NewCache(5 * time.Minute)
	config.MapCache = mapCache

	// Explore cache init
	exploreCache := pkc.NewCache(5 * time.Minute)
	config.ExploreCache = exploreCache

	// REPL loop
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := hp.CleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		// Split input into commandName and params
		commandName, params := words[0], words[1:]

		// Determine command name
		commands := cm.GetCommands()
		command, exists := commands[commandName]

		if exists {
			err := command.Callback(config, params)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println()
			fmt.Println("Unknown command")
			fmt.Println()
		}
	}
}
