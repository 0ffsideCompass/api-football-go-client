package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	timezoneEndpoint = "timezone"
)

// Timezone hits the /timezone endpoint. It returns the list of timezones that
// can be used in the fixtures endpoints. This endpoint takes no parameters.
func (c *Client) Timezone() (*models.TimezoneResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, timezoneEndpoint)
	body, err := c.get(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("error getting timezone: %w", err)
	}

	var resp models.TimezoneResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling timezone response: %w",
			err,
		)
	}

	return &resp, nil
}
