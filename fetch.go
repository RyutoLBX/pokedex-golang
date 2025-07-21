package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Failed to get from %s with status code: %d\nbody:%s\n", url, res.StatusCode, body)
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return body, nil
}
