package models

// SidelinedResponse is the response from the /sidelined endpoint
type SidelinedResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []struct {
		Type  string `json:"type"`
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"response"`
}
