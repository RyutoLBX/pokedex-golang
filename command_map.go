package main

import (
	"encoding/json"
	"fmt"
)

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	if c.Next == nil {
		fmt.Println("You are on the last page of the map!")
		return nil
	}

	locationAreaJson, err := fetch(*c.Next)
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
