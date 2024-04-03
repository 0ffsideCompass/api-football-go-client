package models

// TeamsResponse is the response from the /teams endpoint
type TeamsResponse struct {
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
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
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
			Played struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"played"`
			Wins struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"wins"`
			Draws struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"draws"`
			Loses struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"loses"`
		} `json:"fixtures"`
		Goals struct {
			For struct {
				Total struct {
					Home  int `json:"home"`
					Away  int `json:"away"`
					Total int `json:"total"`
				} `json:"total"`
				Average struct {
					Home  string `json:"home"`
					Away  string `json:"away"`
					Total string `json:"total"`
				} `json:"average"`
				Minute struct {
					Zero15 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"0-15"`
					One630 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"16-30"`
					Three145 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"31-45"`
					Four660 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"46-60"`
					Six175 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"61-75"`
					Seven690 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"76-90"`
					Nine1105 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"91-105"`
					One06120 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"106-120"`
				} `json:"minute"`
			} `json:"for"`
			Against struct {
				Total struct {
					Home  int `json:"home"`
					Away  int `json:"away"`
					Total int `json:"total"`
				} `json:"total"`
				Average struct {
					Home  string `json:"home"`
					Away  string `json:"away"`
					Total string `json:"total"`
				} `json:"average"`
				Minute struct {
					Zero15 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"0-15"`
					One630 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"16-30"`
					Three145 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"31-45"`
					Four660 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"46-60"`
					Six175 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"61-75"`
					Seven690 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"76-90"`
					Nine1105 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"91-105"`
					One06120 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"106-120"`
				} `json:"minute"`
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
		CleanSheet struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"clean_sheet"`
		FailedToScore struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"failed_to_score"`
		Penalty struct {
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
			Yellow struct {
				Zero15 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"0-15"`
				One630 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"16-30"`
				Three145 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"31-45"`
				Four660 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"46-60"`
				Six175 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"61-75"`
				Seven690 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"76-90"`
				Nine1105 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"91-105"`
				One06120 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"106-120"`
			} `json:"yellow"`
			Red struct {
				Zero15 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"0-15"`
				One630 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"16-30"`
				Three145 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"31-45"`
				Four660 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"46-60"`
				Six175 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"61-75"`
				Seven690 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"76-90"`
				Nine1105 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"91-105"`
				One06120 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"106-120"`
			} `json:"red"`
		} `json:"cards"`
	} `json:"response"`
}
