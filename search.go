package urlscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/existentiality/urlscan/structs"
)

func (c Client) Search(query string, size int) (*structs.SearchResult, error) {
	resp, err := c.makeRequest("GET", "/search", RequestData{
		query: map[string]string{
			"q":    query,
			"size": fmt.Sprint(size),
		},
	})

	if err != nil {
		return nil, err
	}

	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseData := *new(structs.SearchResult)
	json.Unmarshal(responseBody, &responseData)

	return &responseData, nil
}
