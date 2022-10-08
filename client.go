package urlscan

import (
	"bytes"
	"fmt"
	"net/http"
)

const Version = "0.0.0"
const BaseUrl = "https://urlscan.io/api/v1"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) Client {
	return Client{
		apiKey,
	}
}

type RequestData struct {
	body  []byte
	query map[string]string
}

func (c Client) makeRequest(method string, path string, data RequestData) (*http.Response, error) {
	bodyReader := bytes.NewReader(data.body)

	// Setup url query
	query := ""
	for key, value := range data.query {
		if len(query) < 1 {
			query = fmt.Sprintf(`?%s=%s`, key, value)
		} else {
			query = fmt.Sprintf(`%s&%s=%s`, query, key, value)
		}
	}

	req, err := http.NewRequest(method, fmt.Sprintf(`%s%s%s`, BaseUrl, path, query), bodyReader)

	if err != nil {
		return nil, err
	}

	req.Header.Set("API-Key", c.apiKey)
	req.Header.Set("User-Agent", fmt.Sprintf("https://github.com/existentiality/urlscan, %s", Version))

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
