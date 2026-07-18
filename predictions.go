package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	predictionsEndpoint = "predictions"
)

// Predictions hits the /predictions endpoint. It returns predictions about a
// fixture, along with both teams' form, comparative statistics, and their
// head-to-head history.
/*
	- fixture: (Type: integer) (Required)
	  The ID of the fixture. Value format: 198772
*/
func (c *Client) Predictions(
	params map[string]any,
) (*models.PredictionsResponse, error) {
	// Validate the parameters
	if err := requireIntParams(params, "fixture"); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, predictionsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting predictions: %w", err)
	}

	var resp models.PredictionsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling predictions response: %w",
			err,
		)
	}

	return &resp, nil
}
