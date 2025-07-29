package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) GetLocationArea(offset, limit int) (LocationAreaShallow, error) {
	url := FormatLocationAreaURL(offset, limit)

	// Make new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaShallow{}, err
	}

	// Send request and get response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaShallow{}, err
	}
	defer resp.Body.Close()

	// Read data from response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaShallow{}, err
	}

	// Read data from response
	locationsResp := LocationAreaShallow{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationAreaShallow{}, err
	}

	return locationsResp, nil
}
