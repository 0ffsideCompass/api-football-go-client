package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	countriesEndpoint = "countries"
)

// Countries hits the /countries endpoint. All parameters are optional and can
// be combined.
/*
	- name: (Type: string)
	  The name of the country. Value format: England
	- code: (Type: string) (2 characters)
	  The Alpha code of the country. Value format: GB
	- search: (Type: string) (should be 3 characters or more)
	  The name of the country. Value format: England
*/
func (c *Client) Countries(
	params map[string]any,
) (*models.CountriesResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, countriesEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting countries: %w", err)
	}

	var resp models.CountriesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling countries response: %w",
			err,
		)
	}

	return &resp, nil
}
