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
	params map[string]any,
) (*models.SidelinedResponse, error) {
	// Validate the parameters
	if err := requireOneIntParam(params, "coach", "player"); err != nil {
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
