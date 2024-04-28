package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	leaguesEndpoint       = "leagues"
	leagueSeasonsEndpoint = "leagueSeasons"
)

// Leagues returns all leagues
/*
	- id (Type: integer)
	  The ID of the league. Value format: 2
	- name (Type: string)
	  The name of the league. Value format: Premier League
	-country (Type: string)
	  The name of the country. Value format: England
	- season (Type: integer) (should be 4 digits)
	  The year of the season. Value format: 2020
	- team (Type: integer)
	  The ID of the team. Value format: 85
	- type (Type: string)
	  The type of the league. Value format: League
	- current (Type: string) Enum: (true, false)
	  Return the list of active seasons or the last one of each competition. Value format: "true"
	- search (Type: string) (should be 3 characters or more)
	  The name of the league. Value format: Premier League
	- last (Type: integer) (<= 2 digis)
	  The X last leagues/cups added in the API
	- code (Type: string) (2 characters)
	  The code of the country. Value format: FR
*/
func (c *Client) Leagues(
	params map[string]interface{},
) (*models.LeaguesResponse, error) {

	endpointURL := fmt.Sprintf("%s%s", c.Domain, leaguesEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting leagues response: %w", err)
	}

	var resp models.LeaguesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling leagues response: %w",
			err,
		)
	}
	return &resp, nil
}
