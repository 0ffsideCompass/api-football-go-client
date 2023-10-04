package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

// Transfers hits the /transfers endpoint
/*
	- player (Type: integer)
	  The ID of the player. Value format: 85
	- team (Type: integer)
	  The ID of the team. Value format: 85
	At least one of the following parameters must be passed
*/
func (c *Client) Transfers(
	params map[string]interface{},
) (*models.TransfersResponse, error) {
	// Validate the parameters
	if err := validateTransfersParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, trophiesEndpoint)

	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting transfers: %w", err)
	}

	var resp models.TransfersResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling transfers response: %w",
			err,
		)
	}

	return &resp, nil
}

func validateTransfersParams(params map[string]interface{}) error {
	if _, hasCoach := params["team"]; !hasCoach {
		if _, hasPlayer := params["player"]; !hasPlayer {
			return fmt.Errorf("at least one of 'team' or 'player' must be provided")
		}
	}

	for key, value := range params {
		if key == "team" || key == "player" {
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
