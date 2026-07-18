package models

// TimezoneResponse is the response from the /timezone endpoint.
type TimezoneResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []string   `json:"response"`
}

// Country is a single country entry as returned by the /countries and
// /teams/countries endpoints.
type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

// CountriesResponse is the response from the /countries and /teams/countries
// endpoints.
type CountriesResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []Country  `json:"response"`
}

// SeasonsResponse is the response from the /leagues/seasons and /teams/seasons
// endpoints. Seasons are 4-digit years.
type SeasonsResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []int      `json:"response"`
}
