package urlscan

import (
	"encoding/json"
	"fmt"
	"io"

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
		errorMessage := "Unable to make new HTTP request"
		err_Make_Req_Formatted := fmt.Errorf("Error is : %s", errorMessage)
		return nil, err_Make_Req_Formatted
	}

	responseBody, _ := io.ReadAll(resp.Body)
	responseData := *new(structs.SearchResult)
	json.Unmarshal(responseBody, &responseData)

	return &responseData, nil
}
