package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	coachsEndpoint = "coachs"
)

// Coachs hits the /coachs endpoint
/*
	- team: (Type: integer)
	  The ID of the team. Value format: 85
	- id: (Type: integer)
	  The ID of the coach. Value format: 85
	- search: (Type: string)
	  The name of the coach. Value format: "John Doe" must be >= 3 characters
	one of the following parameters must be passed
*/
func (c *Client) Coachs(
	params map[string]interface{},
) (*models.Coachs, error) {
	// Validate the parameters
	if err := validateCoachsParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, coachsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting coachs: %w", err)
	}

	var resp models.Coachs
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling coachs response: %w",
			err,
		)
	}

	return &resp, nil
}

// validateCoachsParams validates the parameters for the Coachs endpoint.
func validateCoachsParams(params map[string]interface{}) error {
	// At least one of these parameters must be passed
	if _, hasTeam := params["team"]; !hasTeam {
		if _, hasID := params["id"]; !hasID {
			if _, hasSearch := params["search"]; !hasSearch {
				return fmt.Errorf("at least one of 'team', 'id', or 'search' must be provided")
			}
		}
	}

	// Validate 'team' parameter
	if teamID, ok := params["team"].(int); ok {
		if teamID != int(teamID) {
			return fmt.Errorf("'team' must be an integer")
		}
	}

	// Validate 'id' parameter
	if coachID, ok := params["id"].(int); ok {
		if coachID != int(coachID) {
			return fmt.Errorf("'id' must be an integer")
		}
	}

	// Validate 'search' parameter
	if search, ok := params["search"].(string); ok {
		if len(search) < 3 {
			return fmt.Errorf("'search' must be at least 3 characters long")
		}
	}

	return nil
}
