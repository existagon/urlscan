package urlscan

import (
	"encoding/json"
	"fmt"
	"io"

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
	jsonBody, errMarshalBody := json.Marshal(requestBody)

	if errMarshalBody != nil {
		errorMessage := "Unable to create a new JSON Body from the Request Body"
		err_MarshalReq_Body := fmt.Errorf("Error : %s", errorMessage)
		return nil, err_MarshalReq_Body
	}

	resp, err_Make_Req := c.makeRequest("POST", "/scan", RequestData{
		body: jsonBody,
	})

	if err_Make_Req != nil {
		errorMessage := "Unable to make new HTTP request with Scanned Request Body"
		err_Make_Req_Formatted := fmt.Errorf("Error is : %s", errorMessage)
		return nil, err_Make_Req_Formatted
	}

	if resp.StatusCode != 200 {
		errorMessage, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(fmt.Sprintf("Error Scanning URL: %s: %s", resp.Status, errorMessage))
	}

	responseBody, _ := io.ReadAll(resp.Body)

	var responseData structs.ScanResponse = *new(structs.ScanResponse)
	err_UnMarshall := json.Unmarshal(responseBody, &responseData)

	if err_UnMarshall != nil {
		errorMessage := "Unable to Unmarshall Response Body into ResponseData Struct"
		err_Unmarshall_Formatted := fmt.Errorf("Error is : %s", errorMessage)
		return nil, err_Unmarshall_Formatted
	}

	return &responseData, nil
}
