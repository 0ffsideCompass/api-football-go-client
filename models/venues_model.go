package models

// VenuesResponse is the response from the /venues endpoint
type VenuesResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		City     string `json:"city"`
		Country  string `json:"country"`
		Capacity int    `json:"capacity"`
		Surface  string `json:"surface"`
		Image    string `json:"image"`
	} `json:"response"`
}
