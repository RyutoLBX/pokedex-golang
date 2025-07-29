package pokeapi

import "fmt"

func FormatLocationAreaURL(offset, limit int) string {
	return fmt.Sprintf("%s/location-area/?offset=%d&limit=%d", baseURL, offset, limit)
}

func FormatExploreURL(locationName string) string {
	return fmt.Sprintf("%s/location-area/%s", baseURL, locationName)
}
