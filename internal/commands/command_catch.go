package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

func CommandCatch(config *Config, params []string) error {
	fmt.Println()
	defer fmt.Println()

	if len(params) == 0 {
		return errors.New("no pokemon name provided")
	}

	locationName := params[0]
	data, err := config.PokemonCache.Get(pokeapi.FormatPokemonURL(locationName), pokeapi.Fetch)
	if err != nil {
		return err
	}

	pokemon := pokeapi.Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return err
	}

	randomFloat := rand.Float64()
	threshold := CatchThresholdFunction(pokemon.BaseExperience)
	if randomFloat > threshold {
		// Catch successful
		fmt.Printf("%s successfully caught!\n", pokemon.Name)
		config.Pokedex.AddToPokedex(pokemon)
	} else {
		// Catch failed
		fmt.Printf("Failed to catch %s...\n", pokemon.Name)
	}

	return nil
}

func CatchThresholdFunction(baseExp int) float64 {
	return 0.8 - 0.65*(math.Exp(-(math.Pow(float64(baseExp)/200.0, 2))/2.0))
}
