package pokeapi

import (
	"io"
	"net/http"
	"time"
)

func Fetch(url string) ([]byte, error) {
	client := NewClient(5 * time.Second)

	// Make new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send request and get response
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read data from response
	val, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return val, nil
}
