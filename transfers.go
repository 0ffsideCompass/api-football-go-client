package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	transfersEndpoint = "transfers"
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
	params map[string]any,
) (*models.TransfersResponse, error) {
	// Validate the parameters
	if err := requireOneIntParam(params, "team", "player"); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, transfersEndpoint)

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
