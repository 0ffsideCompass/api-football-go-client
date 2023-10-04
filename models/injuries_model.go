package models

import "time"

// InjuriesResponse is the response from the /injuries endpoints
type InjuriesResponse struct {
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
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Photo  string `json:"photo"`
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"player"`
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Fixture struct {
			ID        int       `json:"id"`
			Timezone  string    `json:"timezone"`
			Date      time.Time `json:"date"`
			Timestamp int       `json:"timestamp"`
		} `json:"fixture"`
		League struct {
			ID      int         `json:"id"`
			Season  int         `json:"season"`
			Name    string      `json:"name"`
			Country string      `json:"country"`
			Logo    string      `json:"logo"`
			Flag    interface{} `json:"flag"`
		} `json:"league"`
	} `json:"response"`
}
