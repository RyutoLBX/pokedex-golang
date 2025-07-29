package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	cm "github.com/RyutoLBX/pokedexcli/internal/commands"
	hp "github.com/RyutoLBX/pokedexcli/internal/helpers"
	"github.com/RyutoLBX/pokedexcli/internal/pokecache"
)

// StartRepl initiates the PokedexCLI REPL loop.
func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &cm.Config{Next: nil, Previous: nil}
	config.Next = hp.Ptr("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")

	mapCache := pokecache.NewCache(5 * time.Minute)
	config.MapCache = mapCache

	exploreCache := pokecache.NewCache(5 * time.Minute)
	config.ExploreCache = exploreCache

	// REPL loop
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := hp.CleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName, params := words[0], words[1:]

		// Handle command (first word)
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
