package models

// PredictionTeam holds a team's form and league statistics as returned by the
// /predictions endpoint.
type PredictionTeam struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	Last5 struct {
		Form  string `json:"form"`
		Att   string `json:"att"`
		Def   string `json:"def"`
		Goals struct {
			For struct {
				Total   int     `json:"total"`
				Average float64 `json:"average"`
			} `json:"for"`
			Against struct {
				Total   int     `json:"total"`
				Average float64 `json:"average"`
			} `json:"against"`
		} `json:"goals"`
	} `json:"last_5"`
	League struct {
		Form     string `json:"form"`
		Fixtures struct {
			Played HomeAwayTotal `json:"played"`
			Wins   HomeAwayTotal `json:"wins"`
			Draws  HomeAwayTotal `json:"draws"`
			Loses  HomeAwayTotal `json:"loses"`
		} `json:"fixtures"`
		Goals struct {
			For struct {
				Total   HomeAwayTotal  `json:"total"`
				Average HomeAwayString `json:"average"`
			} `json:"for"`
			Against struct {
				Total   HomeAwayTotal  `json:"total"`
				Average HomeAwayString `json:"average"`
			} `json:"against"`
		} `json:"goals"`
		Biggest struct {
			Streak struct {
				Wins  int `json:"wins"`
				Draws int `json:"draws"`
				Loses int `json:"loses"`
			} `json:"streak"`
			Wins  HomeAwayString `json:"wins"`
			Loses HomeAwayString `json:"loses"`
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
	} `json:"league"`
}

// PredictionsResponse is the response from the /predictions endpoint.
type PredictionsResponse struct {
	Get        string     `json:"get"`
	Parameters any        `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Pagination `json:"paging"`
	Response   []struct {
		Predictions struct {
			Winner struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Comment string `json:"comment"`
			} `json:"winner"`
			WinOrDraw bool   `json:"win_or_draw"`
			UnderOver string `json:"under_over"`
			Goals     struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"goals"`
			Advice  string `json:"advice"`
			Percent struct {
				Home string `json:"home"`
				Draw string `json:"draw"`
				Away string `json:"away"`
			} `json:"percent"`
		} `json:"predictions"`
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
		} `json:"league"`
		Teams struct {
			Home PredictionTeam `json:"home"`
			Away PredictionTeam `json:"away"`
		} `json:"teams"`
		Comparison struct {
			Form                HomeAwayString `json:"form"`
			Att                 HomeAwayString `json:"att"`
			Def                 HomeAwayString `json:"def"`
			PoissonDistribution HomeAwayString `json:"poisson_distribution"`
			H2H                 HomeAwayString `json:"h2h"`
			Goals               HomeAwayString `json:"goals"`
			Total               HomeAwayString `json:"total"`
		} `json:"comparison"`
		H2H []FixtureResp `json:"h2h"`
	} `json:"response"`
}
