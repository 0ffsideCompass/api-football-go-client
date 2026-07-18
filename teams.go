package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	teamsEndpoint           = "teams"
	teamsStatisticsEndpoint = "teams/statistics"
	teamsSeasonsEndpoint    = "teams/seasons"
	teamsCountriesEndpoint  = "teams/countries"
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
	params map[string]any,
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
	params map[string]any,
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

// TeamsSeasons hits the /teams/seasons endpoint. It returns the list of
// seasons available for a team as 4-digit years.
/*
	- team: (Type: integer) (Required)
	  The ID of the team. Value format: 33
*/
func (c *Client) TeamsSeasons(
	params map[string]any,
) (*models.SeasonsResponse, error) {
	// Validate the parameters
	if err := requireIntParams(params, "team"); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, teamsSeasonsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting teams seasons: %w", err)
	}

	var resp models.SeasonsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling teams seasons response: %w",
			err,
		)
	}

	return &resp, nil
}

// TeamsCountries hits the /teams/countries endpoint. It returns the list of
// countries available for the teams endpoint. This endpoint takes no
// parameters.
func (c *Client) TeamsCountries() (*models.CountriesResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, teamsCountriesEndpoint)
	body, err := c.get(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("error getting teams countries: %w", err)
	}

	var resp models.CountriesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling teams countries response: %w",
			err,
		)
	}

	return &resp, nil
}
