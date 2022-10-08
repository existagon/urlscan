package structs

// Used Internally, JSON Data submitted to scan API endpoint
type InternalScanPostData struct {
	// The URL to scan
	URL string `json:"url"`
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

// JSON Data Response from scan API endpoint
type ScanResponse struct {
	// A message detailing submission success/failure
	Message string `json:"message"`
	// The scan's unique ID
	UUID string `json:"uuid"`
	// A URL for the scan's result page
	Result string `json:"result"`
	// A URL for the scan's result API endpoint
	API string `json:"api"`
	// The visibility level of the scan. Either "public", "unlisted", or "private"
	Visibility string `json:"visibility"`
	// Scan options, currently only user agent
	Options struct {
		UserAgent string `json:"useragent"`
	} `json:"options"`
	// The country the url was scanned from
	Country string `json:"country"`
}
