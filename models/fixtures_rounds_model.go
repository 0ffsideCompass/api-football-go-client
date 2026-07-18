package models

import "encoding/json"

// FixturesRoundsResponse is the response from the /fixtures/rounds endpoint.
type FixturesRoundsResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []Round    `json:"response"`
}

// Round is a single round entry from the /fixtures/rounds endpoint. The API
// returns plain strings by default, and objects that include the round dates
// when the 'dates' parameter is set to true; both shapes decode into this type.
type Round struct {
	Round string   `json:"round"`
	Dates []string `json:"dates"`
}

// UnmarshalJSON accepts either a JSON string or a {round, dates} object.
func (r *Round) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		return json.Unmarshal(data, &r.Round)
	}
	type roundAlias Round
	var a roundAlias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	*r = Round(a)
	return nil
}
