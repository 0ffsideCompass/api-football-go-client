package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	sidelinedEndpoint = "sidelined"
)

// Sidelined hits the /sidelined endpoint
/*
  - player: (Type: integer)
  The ID of the player. Value format: 85
  -coach: (Type: integer)
  The ID of the coach. Value format: 85

  At least one of the following parameters must be passed
*/

func (c *Client) Sidelined(
	params map[string]interface{},
) (*models.SidelinedResponse, error) {
	// Validate the parameters
	if err := validateSidelinedParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, sidelinedEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)

	if err != nil {
		return nil, fmt.Errorf("error getting sidelined: %w", err)
	}

	var resp models.SidelinedResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling sidelined response: %w",
			err,
		)
	}

	return &resp, nil
}

func validateSidelinedParams(params map[string]interface{}) error {
	if _, hasCoach := params["coach"]; !hasCoach {
		if _, hasPlayer := params["player"]; !hasPlayer {
			return fmt.Errorf("at least one of 'coach' or 'player' must be provided")
		}
	}

	for key, value := range params {
		if key == "coach" || key == "player" {
			val, ok := value.(float64)
			if !ok {
				return fmt.Errorf("%s must be an integer", key)
			}

			if val != float64(int(val)) {
				return fmt.Errorf("%s must be an integer", key)
			}
		}
	}
	return nil
}
