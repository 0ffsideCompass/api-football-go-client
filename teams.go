package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	teamsEndpoint           = "teams"
	teamsStatisticsEndpoint = "teams/statistics"
)

// Teams hits the /teams endpoint
/*
	-id: (Type: integer)
	  The ID of the team. Value format: 85
	- name: (Type: string)
	  The name of the team. Value format: Arsenal
	- league: (Type: integer)
	  The ID of the league. Value format: 2
	- season: (Type: integer) (should be 4 digits)
	  The year of the season. Value format: 2020
	- country: (Type: string)
	  The name of the country. Value format: England
	- code: (Type: string) (3 characters)
	  The code of the team. Value format: ARS
	- venue: (Type: integer)
	  The ID of the venue. Value format: 18
	- search: (Type: string) (should be 3 characters or more)
	  The name of the team. Value format: Arsenal

	At least one of the following parameters must be passed
*/
func (c *Client) Teams(
	params map[string]interface{},
) (*models.TeamsResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, teamsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting teams: %w", err)
	}

	var resp models.TeamsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling teams response: %w",
			err,
		)
	}

	return &resp, nil
}

// TeamsStatistics hits the /teams/statistics endpoint
/*
	- league (Type: integer) (required)
	  The ID of the league. Value format: 2
	- season (Type: integer) (required) (should be 4 digits)
	  The year of the season. Value format: 2020
	- team (Type: integer) (required)
	  The ID of the team. Value format: 85
	- date (Type: string) (format: YYYY-MM-DD)
	  The limit date. Value format: 2020-12-01
*/
func (c *Client) TeamsStatistics(
	params map[string]interface{},
) (*models.TeamsStatisticsResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, teamsStatisticsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting teams statistics: %w", err)
	}

	var resp models.TeamsStatisticsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling teams statistics response: %w",
			err,
		)
	}

	return &resp, nil
}
