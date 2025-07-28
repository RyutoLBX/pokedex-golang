package commands

import (
	"fmt"
	"time"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

// Prints out a list of area names
// By default gives 20 pages at a time
// Will stop at last page
func CommandMap(config *Config) error {
	fmt.Println()
	defer fmt.Println()
	if config.Next == nil {
		fmt.Println("You are on the last page of the map!")

		return nil
	}

	client := pokeapi.NewClient(5 * time.Second)

	locationAreas, err := client.ListLocationArea(config.Next)
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
