package commands

import (
	"fmt"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

func CommandInspect(config *Config, params []string) error {
	fmt.Println()
	defer fmt.Println()

	if len(params) == 0 {
		return pokeapi.ErrNoParams("pokemon")
	}

	pokemonName := params[0]

	pokemonData, exists := config.Pokedex.Pokemons[pokemonName]

	if exists {
		fmt.Println("Name:", pokemonData.Name)
		fmt.Println("Height:", pokemonData.Height)
		fmt.Println("Weight:", pokemonData.Weight)
		fmt.Println("Stats:")
		for _, pokemonStat := range pokemonData.Stats {
			fmt.Printf("  - %s: %d\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemonData.Types {
			fmt.Println("  -", pokemonType.Type.Name)
		}
	} else {
		fmt.Printf("You have not caught %s yet!\n", pokemonName)
	}

	return nil
}
