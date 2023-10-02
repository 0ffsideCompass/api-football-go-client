package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	standingsEndpoint = "standings"
)

// Standings hits the /standings endpoint
/*
	- season: (Type: integer) (Required) 4 characters YYYY
	  The season of the standings. Value format: 2019
	- league: (Type: integer) (Required)
	  The ID of the league. Value format: 43
	- team: (Type: integer)
	  The ID of the team. Value format: 85
*/
func (c *Client) Standings(
	params map[string]interface{},
) (*models.StandingsResponse, error) {

	err := c.standingsValidations(params)
	if err != nil {
		return nil, fmt.Errorf("error validating standings params: %w", err)
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, standingsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		))
	if err != nil {
		return nil, fmt.Errorf("error getting standings: %w", err)
	}

	var resp models.StandingsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling standings response: %w",
			err,
		)
	}
	return &resp, nil
}

func (c *Client) standingsValidations(
	params map[string]interface{},
) error {
	// Check if season is provided in params
	season, seasonExists := params["season"]
	if !seasonExists {
		return errors.New("season is required in the arguments")
	}

	// Check if season is provided in params
	_, leagueExists := params["league"]
	if !leagueExists {
		return errors.New("league is required in the arguments")
	}

	// Ensure that season is a valid integer
	seasonInt, ok := season.(int)
	if !ok {
		return errors.New("season must be an integer")
	}

	// Check if season is exactly 4 digits
	seasonStr := strconv.Itoa(seasonInt)
	if len(seasonStr) != 4 {
		return errors.New("season must be a 4-digit integer")
	}

	return nil
}
