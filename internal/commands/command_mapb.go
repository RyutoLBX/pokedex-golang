package commands

import (
	"encoding/json"
	"fmt"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

// CommandMapb prints out a list of area names and goes backwards.
// By default gives 20 pages at a time.
// Will stop at first page.
func CommandMapb(config *Config, _ []string) error {
	fmt.Println()
	defer fmt.Println()

	if config.Previous == nil {
		fmt.Println("You are on the first page of the map!")
		return nil
	}

	data, err := config.MapCache.Get(*config.Previous, pokeapi.Fetch)
	if err != nil {
		return err
	}

	locationAreas := pokeapi.LocationAreaShallow{}
	err = json.Unmarshal(data, &locationAreas)
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
