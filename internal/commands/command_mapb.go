package commands

import (
	"fmt"
	"time"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

func CommandMapb(c *Config) error {
	if c.Previous == nil {
		fmt.Println("You are on the first page of the map!")
		return nil
	}

	client := pokeapi.NewClient(5 * time.Second)

	locationAreas, err := client.ListLocationArea(c.Previous)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous

	return nil
}
