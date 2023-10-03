package models

// TrophiesResponse is the response from the /trophies endpointf
type TrophiesResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		League  string `json:"league"`
		Country string `json:"country"`
		Season  string `json:"season"`
		Place   string `json:"place"`
	} `json:"response"`
}
