package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	venuesEndpoint = "venues"
)

// Venues hits the /venues endpoint
/*
At least one of the following parameters must be passed:

- id: (Type: integer)
  The ID of the league. Value format: 4
- name: (Type: string)
  The name of the league. Value format: "OAKA Stadium"
- country: (Type: string)
  The name of the country. Value format: "Greece"
- city: (Type: string)
  The name of the city. Value format: "Athens"
- search (Type: string)
  The name of the venue. Value format: "OAKA Stadium" should be => 3 characters
*/
func (c *Client) Venues(
	params map[string]interface{},
) (*models.VenuesResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, venuesEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting venues: %w", err)
	}

	var resp models.VenuesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling venues response: %w",
			err,
		)
	}

	return &resp, nil
}
