package commands

import (
	"encoding/json"
	"fmt"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

// CommandMap prints out a list of area names and goes forwards.
// By default gives 20 pages at a time.
// Will stop at last page.
func CommandMap(config *Config, _ []string) error {
	fmt.Println()
	defer fmt.Println()

	if config.Next == nil {
		fmt.Println("You are on the last page of the map!")
		return nil
	}

	data, err := config.MapCache.Get(*config.Next, pokeapi.Fetch)
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
