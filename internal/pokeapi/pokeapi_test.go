package pokeapi

import (
	"net/http"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	const baseTime = 1 * time.Second
	const waitTime = baseTime + 1*time.Second
	const testURL = baseURL + "/location-area"

	client := NewClient(baseTime)
	req1, _ := http.NewRequest("GET", testURL, nil)
	req2, _ := http.NewRequest("GET", testURL, nil)

	_, err := client.httpClient.Do(req1)
	if err != nil {
		t.Errorf("expected to handle request")
		return
	}

	time.Sleep(waitTime)
	_, err = client.httpClient.Do(req2)
	if err == nil {
		t.Errorf("expected to timeout client")
		return
	}
}
