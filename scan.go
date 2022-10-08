package urlscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// Used Internally, JSON Data submitted to scan API endpoint
type InternalScanPostData struct {
	// The URL to scan
	Url string `json:"url"`
	// The visibility level of the scan. Either "public", "unlisted", or "private"
	Visibility string `json:"visibility,omitempty"`
	// User-defined tags used to annotate the scan (e.g. "phishing", "malicious"). Maximum 10 tags.
	Tags []string `json:"tags,omitempty"`
	// Override User-Agent for this scan
	CustomAgent string `json:"customagent,omitempty"`
	// Override HTTP referer for this scan
	Referer string `json:"referer,omitempty"`
	// The country to scan the URL from, see [the API] for a list of possible values
	//
	// [the API]: https://urlscan.io/api/v1/availableCountries
	Country string `json:"country,omitempty"`
}

// Optional values for the Scan API
type ScanOptions struct {
	// The visibility level of the scan. Either "public", "unlisted", or "private"
	visibility string
	// User-defined tags used to annotate the scan (e.g. "phishing", "malicious"). Maximum 10 tags.
	tags []string
	// Override User-Agent for this scan
	customAgent string
	// Override HTTP referer header for this scan
	customReferer string
	// The country to scan the URL from, see [the API] for a list of possible values
	//
	// [the API]: https://urlscan.io/api/v1/availableCountries
	country string
}

type InternalScanResponseOptions struct {
	UserAgent string `json:"useragent"`
}

// Used Internally, JSON Data Response from scan API endpoint
type InternalScanPostResponse struct {
	// A message detailing submission success/failure
	Message string `json:"message"`
	// The scan's unique ID
	Uuid string `json:"uuid"`
	// A URL for the scan's result page
	Result string `json:"result"`
	// A URL for the scan's result API endpoint
	Api string `json:"api"`
	// The visibility level of the scan. Either "public", "unlisted", or "private"
	Visibility string `json:"visibility"`
	// Scan options, currently only user agent
	Options InternalScanResponseOptions `json:"options"`
	// The country the url was scanned from
	Country string `json:"country"`
}

type ScanResponseOptions struct {
	userAgent string
}

type ScanResponse struct {
	// A message detailing submission success/failure
	message string
	// The scan's unique ID
	uuid string
	// A URL for the scan's result page
	resultUrl string
	// A URL for the scan's result API endpoint
	apiResultUrl string
	// The visibility level of the scan. Either "public", "unlisted", or "private"
	visibility string
	// Scan options, currently only user agent
	options ScanResponseOptions
	// The country the url was scanned from
	country string
}

// Scan the Specified URL with the set options
func (c Client) Scan(url string, options ScanOptions) (*ScanResponse, error) {
	requestBody := InternalScanPostData{
		Url:         url,
		Visibility:  options.visibility,
		Tags:        options.tags,
		CustomAgent: options.customAgent,
		Referer:     options.customReferer,
		Country:     options.country,
	}

	jsonBody, err := json.Marshal(requestBody)

	if err != nil {
		return nil, err
	}

	resp, err := c.MakeRequest("POST", "/scan", RequestData{
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

	var responseData *InternalScanPostResponse = &InternalScanPostResponse{}
	err = json.Unmarshal(responseBody, responseData)

	if err != nil {
		return nil, err
	}

	scanResponse := ScanResponse{
		message:      responseData.Message,
		uuid:         responseData.Uuid,
		resultUrl:    responseData.Result,
		apiResultUrl: responseData.Api,
		visibility:   responseData.Visibility,
		options:      ScanResponseOptions{userAgent: responseData.Options.UserAgent},
		country:      responseData.Country,
	}

	return &scanResponse, nil
}
