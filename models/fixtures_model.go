package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// FlexibleValue can unmarshal both string and number types from JSON
type FlexibleValue string

func (fv *FlexibleValue) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as string first
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*fv = FlexibleValue(str)
		return nil
	}
	
	// Try to unmarshal as number
	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*fv = FlexibleValue(fmt.Sprintf("%.0f", num))
		return nil
	}
	
	// If both fail, try as integer
	var intVal int
	if err := json.Unmarshal(data, &intVal); err == nil {
		*fv = FlexibleValue(strconv.Itoa(intVal))
		return nil
	}
	
	// If all fail, set to empty string
	*fv = FlexibleValue("")
	return nil
}

// FixturesStatisticsResponse is the response from the /fixtures/statistics endpoint
type FixturesStatisticsResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     Pagination  `json:"paging"`
	Response   []struct {
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Statistics []struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"statistics"`
	} `json:"response"`
}

// FixturesResponse is the response from the /fixtures endpoint
type FixturesResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     Pagination  `json:"paging"`
	Response   []struct {
		Fixture struct {
			ID        int       `json:"id"`
			Referee   string    `json:"referee"`
			Timezone  string    `json:"timezone"`
			Date      time.Time `json:"date"`
			Timestamp int       `json:"timestamp"`
			Periods   struct {
				First  int `json:"first"`
				Second int `json:"second"`
			} `json:"periods"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				City string `json:"city"`
			} `json:"venue"`
			Status struct {
				Long    string `json:"long"`
				Short   string `json:"short"`
				Elapsed int    `json:"elapsed"`
			} `json:"status"`
		} `json:"fixture"`
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
			Round   string `json:"round"`
		} `json:"league"`
		Teams struct {
			Home struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Winner bool   `json:"winner"`
			} `json:"home"`
			Away struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Winner bool   `json:"winner"`
			} `json:"away"`
		} `json:"teams"`
		Goals struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"goals"`
		Score struct {
			Halftime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"halftime"`
			Fulltime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"fulltime"`
			Extratime struct {
				Home *int `json:"home"`
				Away *int `json:"away"`
			} `json:"extratime"`
			Penalty struct {
				Home *int `json:"home"`
				Away *int `json:"away"`
			} `json:"penalty"`
		} `json:"score"`
		Lineups []struct {
			Team struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Colors struct {
					Player struct {
						Primary string `json:"primary"`
						Number  string `json:"number"`
						Border  string `json:"border"`
					} `json:"player"`
					Goalkeeper struct {
						Primary string `json:"primary"`
						Number  string `json:"number"`
						Border  string `json:"border"`
					} `json:"goalkeeper"`
				} `json:"colors"`
			} `json:"team"`
			Formation string `json:"formation"`
			StartXI   []struct {
				Player struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Number int    `json:"number"`
					Pos    string `json:"pos"`
					Grid   string `json:"grid"`
				} `json:"player"`
			} `json:"startXI"`
			Substitutes []struct {
				Player struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Number int    `json:"number"`
					Pos    string `json:"pos"`
					Grid   string `json:"grid"`
				} `json:"player"`
			} `json:"substitutes"`
			Coach struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Photo string `json:"photo"`
			} `json:"coach"`
		} `json:"lineups"`
		Statistics []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			Statistics []struct {
				Type  string        `json:"type"`
				Value FlexibleValue `json:"value"`
			} `json:"statistics"`
		} `json:"statistics"`
		Players []struct {
			Team struct {
				ID     int       `json:"id"`
				Name   string    `json:"name"`
				Logo   string    `json:"logo"`
				Update time.Time `json:"update"`
			} `json:"team"`
			Players []struct {
				Player struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Photo string `json:"photo"`
				} `json:"player"`
				Statistics []struct {
					Games struct {
						Minutes    int    `json:"minutes"`
						Number     int    `json:"number"`
						Position   string `json:"position"`
						Rating     string `json:"rating"`
						Captain    bool   `json:"captain"`
						Substitute bool   `json:"substitute"`
					} `json:"games"`
					Offsides int `json:"offsides"`
					Shots    struct {
						Total int `json:"total"`
						On    int `json:"on"`
					} `json:"shots"`
					Goals struct {
						Total    int `json:"total"`
						Conceded int `json:"conceded"`
						Assists  int `json:"assists"`
						Saves    int `json:"saves"`
					} `json:"goals"`
					Passes struct {
						Total    int    `json:"total"`
						Key      int    `json:"key"`
						Accuracy string `json:"accuracy"`
					} `json:"passes"`
					Tackles struct {
						Total         int `json:"total"`
						Blocks        int `json:"blocks"`
						Interceptions int `json:"interceptions"`
					} `json:"tackles"`
					Duels struct {
						Total int `json:"total"`
						Won   int `json:"won"`
					} `json:"duels"`
					Dribbles struct {
						Attempts int `json:"attempts"`
						Success  int `json:"success"`
						Past     int `json:"past"`
					} `json:"dribbles"`
					Fouls struct {
						Drawn     int `json:"drawn"`
						Committed int `json:"committed"`
					} `json:"fouls"`
					Cards struct {
						Yellow int `json:"yellow"`
						Red    int `json:"red"`
					} `json:"cards"`
					Penalty struct {
						Won       int `json:"won"`
						Committed int `json:"commited"`
						Scored    int `json:"scored"`
						Missed    int `json:"missed"`
						Saved     int `json:"saved"`
					} `json:"penalty"`
				} `json:"statistics"`
			} `json:"players"`
		} `json:"players"`
		Events []struct {
			Time struct {
				Elapsed int           `json:"elapsed"`
				Extra   FlexibleValue `json:"extra"`
			} `json:"time"`
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			Player struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"player"`
			Assist struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"assist"`
			Type     string `json:"type"`
			Detail   string `json:"detail"`
			Comments string `json:"comments"`
		} `json:"events"`
	} `json:"response"`
}

// FixturesEventsResponse is the response from the /fixtures/events endpoint
type FixturesEventsResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     Pagination  `json:"paging"`
	Response   []struct {
		Time struct {
			Elapsed int           `json:"elapsed"`
			Extra   FlexibleValue `json:"extra"`
		} `json:"time"`
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Player struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"player"`
		Assist struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"assist"`
		Type     string `json:"type"`
		Detail   string `json:"detail"`
		Comments string `json:"comments"`
	} `json:"response"`
}

// FixturesLineupsResponse is the response from the /fixtures/lineups endpoint
type FixturesLineupsResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     Pagination  `json:"paging"`
	Response   []struct {
		Team struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Logo   string `json:"logo"`
			Colors string `json:"colors"`
		} `json:"team"`
		Coach struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Photo string `json:"photo"`
		} `json:"coach"`
		Formation string `json:"formation"`
		StartXI   []struct {
			Player struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Number int    `json:"number"`
				Pos    string `json:"pos"`
				Grid   string `json:"grid"`
			} `json:"player"`
		} `json:"startXI"`
		Substitutes []struct {
			Player struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Number int    `json:"number"`
				Pos    string `json:"pos"`
				Grid   string `json:"grid"`
			} `json:"player"`
		} `json:"substitutes"`
	} `json:"response"`
}

// FixturesPlayersResponse is the response from the /fixtures/players endpoint
type FixturesPlayersResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     Pagination  `json:"paging"`
	Response   []struct {
		Team struct {
			ID     int       `json:"id"`
			Name   string    `json:"name"`
			Logo   string    `json:"logo"`
			Update time.Time `json:"update"`
		} `json:"team"`
		Players []struct {
			Player struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Photo string `json:"photo"`
			} `json:"player"`
			Statistics []struct {
				Games struct {
					Minutes    int    `json:"minutes"`
					Number     int    `json:"number"`
					Position   string `json:"position"`
					Rating     string `json:"rating"`
					Captain    bool   `json:"captain"`
					Substitute bool   `json:"substitute"`
				} `json:"games"`
				Offsides int `json:"offsides"`
				Shots    struct {
					Total int `json:"total"`
					On    int `json:"on"`
				} `json:"shots"`
				Goals struct {
					Total    int `json:"total"`
					Conceded int `json:"conceded"`
					Assists  int `json:"assists"`
					Saves    int `json:"saves"`
				} `json:"goals"`
				Passes struct {
					Total    int    `json:"total"`
					Key      int    `json:"key"`
					Accuracy string `json:"accuracy"`
				} `json:"passes"`
				Tackles struct {
					Total         int `json:"total"`
					Blocks        int `json:"blocks"`
					Interceptions int `json:"interceptions"`
				} `json:"tackles"`
				Duels struct {
					Total int `json:"total"`
					Won   int `json:"won"`
				} `json:"duels"`
				Dribbles struct {
					Attempts int `json:"attempts"`
					Success  int `json:"success"`
					Past     int `json:"past"`
				} `json:"dribbles"`
				Fouls struct {
					Drawn     int `json:"drawn"`
					Committed int `json:"committed"`
				} `json:"fouls"`
				Cards struct {
					Yellow int `json:"yellow"`
					Red    int `json:"red"`
				} `json:"cards"`
				Penalty struct {
					Won       int `json:"won"`
					Committed int `json:"commited"`
					Scored    int `json:"scored"`
					Missed    int `json:"missed"`
					Saved     int `json:"saved"`
				} `json:"penalty"`
			} `json:"statistics"`
		} `json:"players"`
	} `json:"response"`
}
