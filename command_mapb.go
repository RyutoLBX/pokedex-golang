package main

import (
	"encoding/json"
	"fmt"
)

func commandMapb(c *config) error {
	if c.Previous == nil {
		fmt.Println("You are on the first page of the map!")
		return nil
	}

	locationAreaJson, err := fetch(*c.Previous)
	if err != nil {
		return err
	}

	var locationAreaInfo locationArea
	if err := json.Unmarshal(locationAreaJson, &locationAreaInfo); err != nil {
		return err
	}

	locationAreas := locationAreaInfo.Results
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}

	c.Next = locationAreaInfo.Next
	c.Previous = locationAreaInfo.Previous

	return nil
}
