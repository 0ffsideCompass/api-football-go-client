// Package models contains the data structures for API responses from the football API.
package models

import "encoding/json"

// FlexString is a string that also accepts JSON numbers and null when
// unmarshalling. API-Football mixes value types in statistics fields — the
// same array can contain "55%", 12, and null. Numbers keep their literal
// form (12 becomes "12"); null becomes the empty string.
type FlexString string

// UnmarshalJSON accepts a JSON string, number, or null.
func (f *FlexString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*f = ""
		return nil
	}
	if len(data) > 0 && data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*f = FlexString(s)
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	*f = FlexString(n.String())
	return nil
}

// MinuteStat represents statistics for a specific minute range
type MinuteStat struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}

// MinuteBreakdown represents statistics broken down by 15-minute intervals
type MinuteBreakdown struct {
	Zero15   MinuteStat `json:"0-15"`
	One630   MinuteStat `json:"16-30"`
	Three145 MinuteStat `json:"31-45"`
	Four660  MinuteStat `json:"46-60"`
	Six175   MinuteStat `json:"61-75"`
	Seven690 MinuteStat `json:"76-90"`
	Nine1105 MinuteStat `json:"91-105"`
	One06120 MinuteStat `json:"106-120"`
}

// HomeAwayTotal represents statistics with home, away, and total values
type HomeAwayTotal struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

// HomeAwayString represents string statistics with home and away values
type HomeAwayString struct {
	Home  string `json:"home"`
	Away  string `json:"away"`
	Total string `json:"total"`
}

// Pagination represents the paging information in API responses
type Pagination struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}
