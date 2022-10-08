package urlscan

import (
	"fmt"
	"os"
	"testing"
)

var apiKey = os.Getenv("URLSCAN_API_KEY")
var client = NewClient(apiKey)

func TestNewClient(t *testing.T) {
	expected := Client{
		apiKey,
	}

	if client != expected {
		t.Fatalf(`Client generated incorrectly: expected %v, got %v`, expected, client)
	}
}

func TestMakeRequest(t *testing.T) {
	_, err := client.MakeRequest("GET", "/", RequestData{
		query: map[string]string{
			"test": "hi",
			"abc":  "123",
		},
		body: []byte("hi there"),
	},
	)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestScan(t *testing.T) {
	resp, err := client.Scan("https://example.com/", ScanOptions{
		visibility: "public",
		country:    "de",
		tags:       []string{"test"},
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(resp)
}
