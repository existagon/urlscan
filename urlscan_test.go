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
	_, err := client.makeRequest("GET", "/", RequestData{
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
		Visibility: "public",
		Country:    "de",
		Tags:       []string{"test", "urlscan-go"},
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if resp.Message != "Submission successful" {
		t.Fatalf("Unexpected response message: %s", resp.Message)
	}
	if resp.Country != "de" {
		t.Fatalf("Unexpected country: %s", resp.Country)
	}
	if resp.Visibility != "public" {
		t.Fatalf("Unexpected visibility: %s", resp.Visibility)
	}
}

func TestResult(t *testing.T) {
	uuid := "41758b93-b25e-4ff9-ae3a-bde068f0d6c0"
	resp, err := client.GetResult(uuid)

	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(resp)

}

func TestSearch(t *testing.T) {
	resp, err := client.Search("domain:example.com", 100)

	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(resp)

	if len(resp.Results) > 100 {
		t.Fatalf("Unexpected size: %d, expected <=100", len(resp.Results))
	}
}
