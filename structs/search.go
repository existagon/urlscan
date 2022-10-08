package structs

import "time"

type SearchResult struct {
	Results []struct {
		Task struct {
			Visibility string    `json:"visibility"`
			Method     string    `json:"method"`
			Domain     string    `json:"domain"`
			ApexDomain string    `json:"apexDomain"`
			Time       time.Time `json:"time"`
			UUID       string    `json:"uuid"`
			URL        string    `json:"url"`
			Tags       []string  `json:"tags,omitempty"`
		} `json:"task,omitempty"`
		Stats struct {
			UniqIPs           int `json:"uniqIPs"`
			UniqCountries     int `json:"uniqCountries"`
			DataLength        int `json:"dataLength"`
			EncodedDataLength int `json:"encodedDataLength"`
			Requests          int `json:"requests"`
		} `json:"stats"`
		Page struct {
			Country      string    `json:"country"`
			Server       string    `json:"server"`
			IP           string    `json:"ip"`
			MimeType     string    `json:"mimeType"`
			Title        string    `json:"title"`
			URL          string    `json:"url"`
			TLSValidDays int       `json:"tlsValidDays"`
			TLSAgeDays   int       `json:"tlsAgeDays"`
			Ptr          string    `json:"ptr,omitempty"`
			TLSValidFrom time.Time `json:"tlsValidFrom"`
			Domain       string    `json:"domain"`
			UmbrellaRank int       `json:"umbrellaRank"`
			ApexDomain   string    `json:"apexDomain"`
			Asnname      string    `json:"asnname"`
			Asn          string    `json:"asn"`
			TLSIssuer    string    `json:"tlsIssuer"`
			Status       string    `json:"status"`
		} `json:"page,omitempty"`
		ID         string        `json:"_id"`
		Sort       []interface{} `json:"sort"`
		Result     string        `json:"result"`
		Screenshot string        `json:"screenshot"`
	} `json:"results"`
	Total   int  `json:"total"`
	Took    int  `json:"took"`
	HasMore bool `json:"has_more"`
}
