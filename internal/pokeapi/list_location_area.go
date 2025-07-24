package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocationArea(pageURL *string) (LocationAreaShallow, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaShallow{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaShallow{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaShallow{}, err
	}

	locationsResp := LocationAreaShallow{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationAreaShallow{}, err
	}

	return locationsResp, nil
}
