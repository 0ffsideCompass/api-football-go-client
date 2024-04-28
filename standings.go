package client

import (
	"encoding/json"
	"fmt"

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
