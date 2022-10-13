package urlscan

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/existentiality/urlscan/structs"
)

// Get the urlscan result for the specified scan UUID
func (c Client) GetResult(uuid string) (*structs.ResultData, error) {
	resp, err := c.makeRequest("GET", fmt.Sprintf("/result/%s/", uuid), RequestData{})

	if err != nil {
		errorMessage := "Can't make new HTTP request new get result"
		errGetRequestFormatted := fmt.Errorf("Error is : %s", errorMessage)
		return nil, errGetRequestFormatted
	}

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		errorMessage := "Unable to read Response Body"
		errBodyReader := fmt.Errorf("Error is : %s", errorMessage)
		return nil, errBodyReader
	}

	responseData := *new(structs.ResultData)
	json.Unmarshal(responseBody, &responseData)

	return &responseData, nil
}
