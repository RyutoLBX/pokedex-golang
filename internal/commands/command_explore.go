package commands

import (
	"encoding/json"
	"fmt"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

// CommandExplore returns a list of Pokemon from a certain area.
// Only supports exactly one parameter containing the area name.
func CommandExplore(config *Config, params []string) error {
	fmt.Println()
	defer fmt.Println()

	if len(params) == 0 {
		return pokeapi.ErrNoParams("area")
	}

	locationName := params[0]
	data, err := config.ExploreCache.Get(pokeapi.FormatExploreURL(locationName), pokeapi.Fetch)
	if err != nil {
		return err
	}

	locationExplores := pokeapi.LocationAreaExplore{}
	err = json.Unmarshal(data, &locationExplores)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationExplores.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
