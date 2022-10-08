package urlscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/existentiality/urlscan/structs"
)

// Get the urlscan result for the specified scan UUID
func (c Client) GetResult(uuid string) (*structs.ResultData, error) {
	resp, err := c.makeRequest("GET", fmt.Sprintf("/result/%s/", uuid), RequestData{})

	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	responseData := *new(structs.ResultData)
	json.Unmarshal(responseBody, &responseData)

	return &responseData, nil
}
