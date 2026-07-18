package models

// TrophiesResponse is the response from the /trophies endpointf
type TrophiesResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []struct {
		League  string `json:"league"`
		Country string `json:"country"`
		Season  string `json:"season"`
		Place   string `json:"place"`
	} `json:"response"`
}
