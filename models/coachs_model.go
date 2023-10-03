package models

// Coachs is the response from the /coachs endpoint
type Coachs struct {
	Get        string      `json:"get"`
	Parameters interface{} `json:"parameters"`
	Errors     interface{} `json:"errors"`
	Results    int         `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
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
		Nationality string      `json:"nationality"`
		Height      interface{} `json:"height"`
		Weight      interface{} `json:"weight"`
		Photo       string      `json:"photo"`
		Team        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Career []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"career"`
	} `json:"response"`
}
