package models

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