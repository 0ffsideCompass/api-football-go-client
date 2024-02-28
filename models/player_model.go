package models

// PlayersSeasonsResponse is the response from the /players/seasons endpoint
type PlayersSeasonsResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []int `json:"response"`
}

// PlayersResponse is the response from the /players endpoint
type PlayersResponse struct {
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
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
			Age       int    `json:"age"`
			Birth     struct {
				Date    string `json:"date"`
				Place   string `json:"place"`
				Country string `json:"country"`
			} `json:"birth"`
			Nationality string `json:"nationality"`
			Height      string `json:"height"`
			Weight      string `json:"weight"`
			Injured     bool   `json:"injured"`
			Photo       string `json:"photo"`
		} `json:"player"`
		Statistics []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			League struct {
				ID      int         `json:"id"`
				Name    string      `json:"name"`
				Country string      `json:"country"`
				Logo    string      `json:"logo"`
				Flag    interface{} `json:"flag"`
				Season  int         `json:"season"`
			} `json:"league"`
			Games struct {
				Appearences int         `json:"appearences"`
				Lineups     int         `json:"lineups"`
				Minutes     int         `json:"minutes"`
				Number      interface{} `json:"number"`
				Position    string      `json:"position"`
				Rating      interface{} `json:"rating"`
				Captain     bool        `json:"captain"`
			} `json:"games"`
			Substitutes struct {
				In    int `json:"in"`
				Out   int `json:"out"`
				Bench int `json:"bench"`
			} `json:"substitutes"`
			Shots struct {
				Total interface{} `json:"total"`
				On    interface{} `json:"on"`
			} `json:"shots"`
			Goals struct {
				Total    int         `json:"total"`
				Conceded interface{} `json:"conceded"`
				Assists  interface{} `json:"assists"`
				Saves    interface{} `json:"saves"`
			} `json:"goals"`
			Passes struct {
				Total    interface{} `json:"total"`
				Key      interface{} `json:"key"`
				Accuracy interface{} `json:"accuracy"`
			} `json:"passes"`
			Tackles struct {
				Total         interface{} `json:"total"`
				Blocks        interface{} `json:"blocks"`
				Interceptions interface{} `json:"interceptions"`
			} `json:"tackles"`
			Duels struct {
				Total interface{} `json:"total"`
				Won   interface{} `json:"won"`
			} `json:"duels"`
			Dribbles struct {
				Attempts interface{} `json:"attempts"`
				Success  interface{} `json:"success"`
				Past     interface{} `json:"past"`
			} `json:"dribbles"`
			Fouls struct {
				Drawn     interface{} `json:"drawn"`
				Committed interface{} `json:"committed"`
			} `json:"fouls"`
			Cards struct {
				Yellow    int `json:"yellow"`
				Yellowred int `json:"yellowred"`
				Red       int `json:"red"`
			} `json:"cards"`
			Penalty struct {
				Won      interface{} `json:"won"`
				Commited interface{} `json:"commited"`
				Scored   interface{} `json:"scored"`
				Missed   interface{} `json:"missed"`
				Saved    interface{} `json:"saved"`
			} `json:"penalty"`
		} `json:"statistics"`
	} `json:"response"`
}

// PlayersSearchResponse is the response from the /players/search endpoint
type PlayersSquadsResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Players []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Age      int    `json:"age"`
			Number   int    `json:"number"`
			Position string `json:"position"`
			Photo    string `json:"photo"`
		} `json:"players"`
	} `json:"response"`
}

// PlayersTopResponse is the response from the /players/topscorers, /players/topassists,
// /players/topyellowcards, /players/topredcards endpoints
type PlayersTopResponse struct {
	Get        string `json:"get"`
	Parameters struct {
		League string `json:"league"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Player struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
			Age       int    `json:"age"`
			Birth     struct {
				Date    string      `json:"date"`
				Place   interface{} `json:"place"`
				Country string      `json:"country"`
			} `json:"birth"`
			Nationality string      `json:"nationality"`
			Height      interface{} `json:"height"`
			Weight      interface{} `json:"weight"`
			Injured     bool        `json:"injured"`
			Photo       string      `json:"photo"`
		} `json:"player"`
		Statistics []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			League struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Country string `json:"country"`
				Logo    string `json:"logo"`
				Flag    string `json:"flag"`
				Season  int    `json:"season"`
			} `json:"league"`
			Games struct {
				Appearences int         `json:"appearences"`
				Lineups     int         `json:"lineups"`
				Minutes     int         `json:"minutes"`
				Number      interface{} `json:"number"`
				Position    string      `json:"position"`
				Rating      interface{} `json:"rating"`
				Captain     bool        `json:"captain"`
			} `json:"games"`
			Substitutes struct {
				In    int `json:"in"`
				Out   int `json:"out"`
				Bench int `json:"bench"`
			} `json:"substitutes"`
			Shots struct {
				Total int `json:"total"`
				On    int `json:"on"`
			} `json:"shots"`
			Goals struct {
				Total    int         `json:"total"`
				Conceded interface{} `json:"conceded"`
				Assists  interface{} `json:"assists"`
				Saves    interface{} `json:"saves"`
			} `json:"goals"`
			Passes struct {
				Total    interface{} `json:"total"`
				Key      interface{} `json:"key"`
				Accuracy interface{} `json:"accuracy"`
			} `json:"passes"`
			Tackles struct {
				Total         interface{} `json:"total"`
				Blocks        interface{} `json:"blocks"`
				Interceptions interface{} `json:"interceptions"`
			} `json:"tackles"`
			Duels struct {
				Total interface{} `json:"total"`
				Won   interface{} `json:"won"`
			} `json:"duels"`
			Dribbles struct {
				Attempts interface{} `json:"attempts"`
				Success  interface{} `json:"success"`
				Past     interface{} `json:"past"`
			} `json:"dribbles"`
			Fouls struct {
				Drawn     interface{} `json:"drawn"`
				Committed interface{} `json:"committed"`
			} `json:"fouls"`
			Cards struct {
				Yellow    int `json:"yellow"`
				Yellowred int `json:"yellowred"`
				Red       int `json:"red"`
			} `json:"cards"`
			Penalty struct {
				Won      interface{} `json:"won"`
				Commited interface{} `json:"commited"`
				Scored   interface{} `json:"scored"`
				Missed   interface{} `json:"missed"`
				Saved    interface{} `json:"saved"`
			} `json:"penalty"`
		} `json:"statistics"`
	} `json:"response"`
}
