package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	cm "github.com/RyutoLBX/pokedexcli/internal/commands"
	hp "github.com/RyutoLBX/pokedexcli/internal/helpers"
	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
	pkc "github.com/RyutoLBX/pokedexcli/internal/pokecache"
)

// StartRepl initiates the PokedexCLI REPL loop.
func StartRepl() {
	// Scanner init
	scanner := bufio.NewScanner(os.Stdin)

	// Config init
	config := &cm.Config{Next: nil, Previous: nil}
	config.Next = hp.Ptr(pokeapi.FormatLocationAreaURL(0, 20))

	// Map cache init
	mapCache := pkc.NewCache(1 * time.Hour)
	config.MapCache = mapCache
	dataArea, err := config.MapCache.Get(pokeapi.FormatLocationAreaURL(0, 10000), pokeapi.Fetch)
	if err != nil {
		fmt.Println("Error:", err)
	}

	locationAreas := pokeapi.LocationAreaShallow{}
	err = json.Unmarshal(dataArea, &locationAreas)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Explore cache init
	exploreCache := pkc.NewCache(1 * time.Hour)
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
			fmt.Println("Error:", pokeapi.ErrUnknownCommand)
			fmt.Println()
		}
	}
}
