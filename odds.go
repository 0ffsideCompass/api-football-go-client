package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	oddsEndpoint           = "odds"
	oddsMappingEndpoint    = "odds/mapping"
	oddsBookmakersEndpoint = "odds/bookmakers"
	oddsBetsEndpoint       = "odds/bets"
	oddsLiveEndpoint       = "odds/live"
	oddsLiveBetsEndpoint   = "odds/live/bets"
)

// Odds hits the /odds endpoint. It returns pre-match odds for fixtures.
/*
	- fixture: (Type: integer)
	  The ID of the fixture. Value format: 326090
	- league: (Type: integer)
	  The ID of the league. Value format: 116
	- season: (Type: integer) (should be 4 digits)
	  The season of the league. Value format: 2020
	- date: (Type: string) (format: YYYY-MM-DD)
	  A valid date. Value format: 2020-05-15
	- timezone: (Type: string)
	  A valid timezone from the Timezone endpoint. Value format: Europe/London
	- page: (Type: integer)
	  Use for pagination. Value format: 1
	- bookmaker: (Type: integer)
	  The ID of the bookmaker. Value format: 6
	- bet: (Type: integer)
	  The ID of the bet. Value format: 38
*/
func (c *Client) Odds(
	params map[string]any,
) (*models.OddsResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, oddsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting odds: %w", err)
	}

	var resp models.OddsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling odds response: %w",
			err,
		)
	}

	return &resp, nil
}

// OddsMapping hits the /odds/mapping endpoint. It returns the list of fixture
// IDs for which pre-match odds are available.
/*
	- page: (Type: integer)
	  Use for pagination. Value format: 1
*/
func (c *Client) OddsMapping(
	params map[string]any,
) (*models.OddsMappingResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, oddsMappingEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting odds mapping: %w", err)
	}

	var resp models.OddsMappingResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling odds mapping response: %w",
			err,
		)
	}

	return &resp, nil
}

// OddsBookmakers hits the /odds/bookmakers endpoint. It returns the list of
// available bookmakers for the pre-match odds endpoint.
/*
	- id: (Type: integer)
	  The ID of the bookmaker. Value format: 6
	- search: (Type: string)
	  The name of the bookmaker. Value format: Bwin
*/
func (c *Client) OddsBookmakers(
	params map[string]any,
) (*models.OddsBookmakersResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, oddsBookmakersEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting odds bookmakers: %w", err)
	}

	var resp models.OddsBookmakersResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling odds bookmakers response: %w",
			err,
		)
	}

	return &resp, nil
}

// OddsBets hits the /odds/bets endpoint. It returns the list of available bet
// types for the pre-match odds endpoint.
/*
	- id: (Type: integer)
	  The ID of the bet. Value format: 5
	- search: (Type: string)
	  The name of the bet. Value format: under
*/
func (c *Client) OddsBets(
	params map[string]any,
) (*models.OddsBetsResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, oddsBetsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting odds bets: %w", err)
	}

	var resp models.OddsBetsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling odds bets response: %w",
			err,
		)
	}

	return &resp, nil
}

// OddsLive hits the /odds/live endpoint. It returns in-play odds for fixtures
// that are currently in progress.
/*
	- fixture: (Type: integer)
	  The ID of the fixture. Value format: 721238
	- league: (Type: integer)
	  The ID of the league. Value format: 30
	- bet: (Type: integer)
	  The ID of the bet. Value format: 20
*/
func (c *Client) OddsLive(
	params map[string]any,
) (*models.OddsLiveResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, oddsLiveEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting odds live: %w", err)
	}

	var resp models.OddsLiveResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling odds live response: %w",
			err,
		)
	}

	return &resp, nil
}

// OddsLiveBets hits the /odds/live/bets endpoint. It returns the list of
// available bet types for the in-play odds endpoint.
/*
	- id: (Type: integer)
	  The ID of the bet. Value format: 1
	- search: (Type: string)
	  The name of the bet. Value format: under
*/
func (c *Client) OddsLiveBets(
	params map[string]any,
) (*models.OddsBetsResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, oddsLiveBetsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting odds live bets: %w", err)
	}

	var resp models.OddsBetsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling odds live bets response: %w",
			err,
		)
	}

	return &resp, nil
}
