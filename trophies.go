package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	trophiesEndpoint = "trophies"
)

// Trophies hits the /trophies endpoint
/*
	- player: (Type: integer)
	  The ID of the player. Value format: 85
	- coach: (Type: integer)
	  The ID of the coach. Value format: 85

	At least one of the following parameters must be passed
*/
func (c *Client) Trophies(
	params map[string]interface{},
) (*models.TrophiesResponse, error) {
	// Validate the parameters
	if err := validateTrophiesParams(params); err != nil {
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
		return nil, fmt.Errorf("error getting trophies: %w", err)
	}

	var resp models.TrophiesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling trophies response: %w",
			err,
		)
	}

	return &resp, nil
}

func validateTrophiesParams(params map[string]interface{}) error {
	if _, hasCoach := params["coach"]; !hasCoach {
		if _, hasPlayer := params["player"]; !hasPlayer {
			return fmt.Errorf("at least one of 'coach' or 'player' must be provided")
		}
	}

	for key, value := range params {
		if key == "coach" || key == "player" {
			val, ok := value.(int)
			if !ok {
				return fmt.Errorf("%s must be an integer", key)
			}

			if val != int(int(val)) {
				return fmt.Errorf("%s must be an integer", key)
			}
		}
	}
	return nil
}
