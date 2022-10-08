package urlscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/existentiality/urlscan/structs"
)

// Optional values for the Scan API
type ScanOptions struct {
	// The visibility level of the scan. Either "public", "unlisted", or "private"
	Visibility string
	// User-defined tags used to annotate the scan (e.g. "phishing", "malicious"). Maximum 10 tags.
	Tags []string
	// Override User-Agent for this scan
	CustomAgent string
	// Override HTTP referer header for this scan
	CustomReferer string
	// The country to scan the URL from, see [the API] for a list of possible values
	//
	// [the API]: https://urlscan.io/api/v1/availableCountries
	Country string
}

// Scan the Specified URL with the set options
func (c Client) Scan(url string, options ScanOptions) (*structs.ScanResponse, error) {
	requestBody := structs.InternalScanPostData{
		URL:         url,
		Visibility:  options.Visibility,
		Tags:        options.Tags,
		CustomAgent: options.CustomAgent,
		Referer:     options.CustomReferer,
		Country:     options.Country,
	}
	jsonBody, err := json.Marshal(requestBody)

	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest("POST", "/scan", RequestData{
		body: jsonBody,
	})

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		errorMessage, _ := ioutil.ReadAll(resp.Body)

		return nil, errors.New(fmt.Sprintf("Error Scanning URL: %s: %s", resp.Status, errorMessage))
	}

	responseBody, _ := ioutil.ReadAll(resp.Body)

	var responseData structs.ScanResponse = *new(structs.ScanResponse)
	err = json.Unmarshal(responseBody, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
