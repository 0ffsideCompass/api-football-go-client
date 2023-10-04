package models

import "time"

// TransfersResponse is the response from the /transfers endpoint
type TransfersResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Player struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"player"`
		Update    time.Time `json:"update"`
		Transfers []struct {
			Date  string `json:"date"`
			Type  string `json:"type"`
			Teams struct {
				In struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Logo string `json:"logo"`
				} `json:"in"`
				Out struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Logo string `json:"logo"`
				} `json:"out"`
			} `json:"teams"`
		} `json:"transfers"`
	} `json:"response"`
}
