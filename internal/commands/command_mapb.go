package commands

import (
	"fmt"
	"time"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

// Prints out a list of area names, but goes backwards
// By default gives 20 pages at a time
// Will stop at first page
func CommandMapb(config *Config) error {
	if config.Previous == nil {
		fmt.Println("You are on the first page of the map!")
		return nil
	}

	client := pokeapi.NewClient(5 * time.Second)

	locationAreas, err := client.ListLocationArea(config.Previous)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous

	return nil
}
