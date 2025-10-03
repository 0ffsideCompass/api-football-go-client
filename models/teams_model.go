package models

// TeamsResponse is the response from the /teams endpoint
type TeamsResponse struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     Pagination  `json:"paging"`
	Response   []struct {
		Team struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Code     string `json:"code"`
			Country  string `json:"country"`
			Founded  int    `json:"founded"`
			National bool   `json:"national"`
			Logo     string `json:"logo"`
		} `json:"team"`
		Venue struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Address  string `json:"address"`
			City     string `json:"city"`
			Capacity int    `json:"capacity"`
			Surface  string `json:"surface"`
			Image    string `json:"image"`
		} `json:"venue"`
	} `json:"response"`
}

// TeamsStatisticsResponse is the response from the /teams/statistics endpoint
type TeamsStatisticsResponse struct {
	Get        string `json:"get"`
	Parameters struct {
		Season string `json:"season"`
		Team   string `json:"team"`
		League string `json:"league"`
	} `json:"parameters"`
	Errors   []interface{} `json:"errors"`
	Results  int           `json:"results"`
	Paging   Pagination    `json:"paging"`
	Response struct {
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
		} `json:"league"`
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Form     string `json:"form"`
		Fixtures struct {
			Played HomeAwayTotal `json:"played"`
			Wins   HomeAwayTotal `json:"wins"`
			Draws  HomeAwayTotal `json:"draws"`
			Loses  HomeAwayTotal `json:"loses"`
		} `json:"fixtures"`
		Goals struct {
			For struct {
				Total   HomeAwayTotal   `json:"total"`
				Average HomeAwayString  `json:"average"`
				Minute  MinuteBreakdown `json:"minute"`
			} `json:"for"`
			Against struct {
				Total   HomeAwayTotal   `json:"total"`
				Average HomeAwayString  `json:"average"`
				Minute  MinuteBreakdown `json:"minute"`
			} `json:"against"`
		} `json:"goals"`
		Biggest struct {
			Streak struct {
				Wins  int `json:"wins"`
				Draws int `json:"draws"`
				Loses int `json:"loses"`
			} `json:"streak"`
			Wins struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"wins"`
			Loses struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"loses"`
			Goals struct {
				For struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"for"`
				Against struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"against"`
			} `json:"goals"`
		} `json:"biggest"`
		CleanSheet    HomeAwayTotal `json:"clean_sheet"`
		FailedToScore HomeAwayTotal `json:"failed_to_score"`
		Penalty       struct {
			Scored struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"scored"`
			Missed struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"missed"`
			Total int `json:"total"`
		} `json:"penalty"`
		Lineups []struct {
			Formation string `json:"formation"`
			Played    int    `json:"played"`
		} `json:"lineups"`
		Cards struct {
			Yellow MinuteBreakdown `json:"yellow"`
			Red    MinuteBreakdown `json:"red"`
		} `json:"cards"`
	} `json:"response"`
}
