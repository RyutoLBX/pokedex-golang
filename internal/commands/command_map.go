package commands

import (
	"fmt"
	"time"

	"github.com/RyutoLBX/pokedexcli/internal/pokeapi"
)

func CommandMap(c *Config) error {
	if c.Next == nil {
		fmt.Println("You are on the last page of the map!")
		return nil
	}

	client := pokeapi.NewClient(5 * time.Second)

	locationAreas, err := client.ListLocationArea(c.Next)
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
