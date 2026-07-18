package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	coachsEndpoint = "coachs"

	// minSearchLength is the minimum length the API accepts for 'search' values.
	minSearchLength = 3
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
	params map[string]any,
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
func validateCoachsParams(params map[string]any) error {
	// At least one of these parameters must be passed
	if _, hasTeam := params["team"]; !hasTeam {
		if _, hasID := params["id"]; !hasID {
			if _, hasSearch := params["search"]; !hasSearch {
				return fmt.Errorf("at least one of 'team', 'id', or 'search' must be provided")
			}
		}
	}

	// Validate 'search' parameter
	if search, ok := params["search"].(string); ok {
		if len(search) < minSearchLength {
			return fmt.Errorf("'search' must be at least %d characters long", minSearchLength)
		}
	}

	return nil
}
