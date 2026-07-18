package models

import "time"

// BetValue is a single selection within a bet. Value is the selection label,
// which the API returns either as a string (e.g. "Over 2.5") or as a number
// (e.g. an exact goals count), so it is left untyped.
type BetValue struct {
	Value any    `json:"value"`
	Odd   string `json:"odd"`
}

// IDName is a simple id/name pair used by the bookmakers and bets endpoints.
type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// OddsResponse is the response from the /odds (pre-match) endpoint.
type OddsResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []struct {
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
		} `json:"league"`
		Fixture struct {
			ID        int       `json:"id"`
			Timezone  string    `json:"timezone"`
			Date      time.Time `json:"date"`
			Timestamp int64     `json:"timestamp"`
		} `json:"fixture"`
		Update     time.Time `json:"update"`
		Bookmakers []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Bets []struct {
				ID     int        `json:"id"`
				Name   string     `json:"name"`
				Values []BetValue `json:"values"`
			} `json:"bets"`
		} `json:"bookmakers"`
	} `json:"response"`
}

// OddsMappingResponse is the response from the /odds/mapping endpoint.
type OddsMappingResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []struct {
		League struct {
			ID     int `json:"id"`
			Season int `json:"season"`
		} `json:"league"`
		Fixture struct {
			ID        int       `json:"id"`
			Date      time.Time `json:"date"`
			Timestamp int64     `json:"timestamp"`
		} `json:"fixture"`
		Update time.Time `json:"update"`
	} `json:"response"`
}

// OddsBookmakersResponse is the response from the /odds/bookmakers endpoint.
type OddsBookmakersResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []IDName   `json:"response"`
}

// OddsBetsResponse is the response from the /odds/bets and /odds/live/bets
// endpoints.
type OddsBetsResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []IDName   `json:"response"`
}

// OddsLiveResponse is the response from the /odds/live (in-play) endpoint.
type OddsLiveResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []struct {
		Fixture struct {
			ID     int `json:"id"`
			Status struct {
				Long    string `json:"long"`
				Elapsed int    `json:"elapsed"`
				Seconds string `json:"seconds"`
			} `json:"status"`
		} `json:"fixture"`
		League struct {
			ID     int `json:"id"`
			Season int `json:"season"`
		} `json:"league"`
		Teams struct {
			Home struct {
				ID    int `json:"id"`
				Goals int `json:"goals"`
			} `json:"home"`
			Away struct {
				ID    int `json:"id"`
				Goals int `json:"goals"`
			} `json:"away"`
		} `json:"teams"`
		Status struct {
			Stopped  bool `json:"stopped"`
			Blocked  bool `json:"blocked"`
			Finished bool `json:"finished"`
		} `json:"status"`
		Update time.Time `json:"update"`
		Odds   []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Values []struct {
				Value     any    `json:"value"`
				Odd       string `json:"odd"`
				Handicap  string `json:"handicap"`
				Main      any    `json:"main"`
				Suspended bool   `json:"suspended"`
			} `json:"values"`
		} `json:"odds"`
	} `json:"response"`
}
