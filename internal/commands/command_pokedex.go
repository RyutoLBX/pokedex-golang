package commands

import "fmt"

func CommandPokedex(config *Config, _ []string) error {
	fmt.Println()
	defer fmt.Println()

	if len(config.Pokedex.Pokemons) == 0 {
		fmt.Println("Your Pokedex is empty!")
		fmt.Println("Try catching a pokemon using the catch command.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.Pokedex.Pokemons {
		fmt.Println("  -", pokemon.Name)
	}
	return nil
}
