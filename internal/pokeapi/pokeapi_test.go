package pokeapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	const timeout = 1 * time.Second
	const responseTimeSlow = timeout + 1*time.Second

	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	}))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(responseTimeSlow) // simulate slow server
		fmt.Fprintln(w, "Hello World.")
	}))
	defer server2.Close()

	client := NewClient(timeout)
	req1, _ := http.NewRequest("GET", server1.URL, nil)
	req2, _ := http.NewRequest("GET", server2.URL, nil)

	_, err := client.httpClient.Do(req1)
	if err != nil {
		t.Errorf("expected to handle request")
		return
	}

	_, err = client.httpClient.Do(req2)
	if err == nil {
		t.Errorf("expected to timeout client")
		return
	}
}
