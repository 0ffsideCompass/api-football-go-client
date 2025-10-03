package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	fixturesByDateEndpoint    = "fixtures?league=%d&season=%d&from=%s&to=%s"
	fixtureHeadToHeadEndpoint = "fixtures/headtohead"
	fixtureRounds             = "fixtures/rounds?league=%d&season=%d" //TODO not important yet
	fixtureEndpoint           = "fixtures"
	fixtureStatisticsEndpoint = "fixtures/statistics"
	fixturesEventsEndpoint    = "fixtures/events"
	fixtureLineupsEndpoint    = "fixtures/lineups"
	fixturePlayerEndpoint     = "fixtures/players"
)

// FixturesLineups returns lineups for a given fixture
/*
	- fixture: (Type: integer)(Required)
  	The ID of the fixture. Value format: 43.
	- team: (Type: integer)
	The ID of the team. Value format: 85.
	- player: (Type: integer)
	The ID of the player. Value format: 85.
*/
func (c *Client) FixturesLineups(
	params map[string]interface{},
) (*models.FixturesLineupsResponse, error) {

	endpointURL := fmt.Sprintf("%s%s", c.Domain, fixturesEventsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting fixtures events: %w", err)
	}

	var resp models.FixturesLineupsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling fixtures lineups response: %w",
			err,
		)
	}
	return &resp, nil
}

// FixturesEvents returns all events for a given fixture
/*
	- fixture: (Type: integer)(Required)
  	The ID of the fixture. Value format: 43.
	- team: (Type: integer)
	The ID of the team. Value format: 85.
	- player: (Type: integer)
	The ID of the player. Value format: 85.
*/
func (c *Client) FixturesEvents(
	params map[string]interface{},
) (*models.FixturesEventsResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, fixturesEventsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting fixtures events: %w", err)
	}

	var resp models.FixturesEventsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling fixtures events response: %w",
			err,
		)
	}
	return &resp, nil
}

// Fixture returns all fixtures for parameters passed
/*
- id: (Type: integer)
  The ID of the fixture. Value format: 43.

- ids: (Type: strings)
  One or more fixture IDs. Value format: "id-id-id".
  Maximum of 20 fixture IDs.

- live: (Type: string)
  Enumerated values:
  * "all": All leagues.
  * "id-id": Specific league IDs (this is a placeholder; handle dynamically based on IDs).

- date: (Type: string)
  A valid date in the format YYYY-MM-DD.

- league: (Type: integer)
  The ID of the league.

- season: (Type: integer) (should be 4 digits)
  The season of the league. Format: 4 characters YYYY.

- team: (Type: integer)
  The ID of the team.

- last: (Type: integer) (<= 2 digis)
  For the X last fixtures. Value should be an integer with <= 2 characters.

- next: (Type: integer) (<= 2 digis)
  For the X next fixtures. Value should be an integer with <= 2 characters.

- from: (Type: string)
  A valid starting date for a date range. Format: YYYY-MM-DD.

- to: (Type: string)
  A valid ending date for a date range. Format: YYYY-MM-DD.

- round: (Type: string)
  The round of the fixture.

- status: (Type: string)
  Enumerated values:
  * "NS": A specific fixture status.
  * "NS-PST-FT": Combined fixture statuses.

- venue: (Type: integer)
  The venue ID of the fixture.

==================================================================================
*/
func (c *Client) Fixture(
	params map[string]interface{},
) (*models.FixturesResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, fixtureEndpoint)

	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting fixtures: %w", err)
	}
	var resp models.FixturesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling fixtures response: %w",
			err,
		)
	}
	return &resp, nil
}

// FixtureHeadToHead returns head to head for two teams
/*
	- h2h: (Type: string)(Required)(format id-id)
	The ids of the teams
	- date (Type: string)(format YYYY-MM-DD)
	The date of the match
	- league (Type: integer)
	The id of the league
	- season (Type: integer)(4 digits)(format YYYY)
	The season of the league
	- last (Type: integer)
	For the X last fixtures.
	- next (Type: integer)
	For the X next fixtures.
	- from (Type: string)(format YYYY-MM-DD)
	A valid starting date for a date range.
	- to (Type: string)(format YYYY-MM-DD)
	A valid ending date for a date range.
	- venue (Type: integer)
	The venue ID of the fixture.
	- status (Type: string)
	Enumerated values:
	* "NS": A specific fixture status.
	* "NS-PST-FT": Combined fixture statuses.
*/
func (c *Client) FixtureHeadToHead(
	params map[string]interface{},
) (*models.FixtureHeadToHeadResp, error) {

	endpointURL := fmt.Sprintf("%s%s", c.Domain, fixtureHeadToHeadEndpoint)

	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, err
	}

	var resp models.FixtureHeadToHeadResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling fixtures head to head response: %w", err)
	}
	return &resp, nil
}

// FixtureByDateAndLeague returns all fixtures for a given date and league
func (c *Client) FixtureByDateAndLeague(
	leagueID,
	season int,
	fromDate time.Time,
	toDate time.Time,
) (*models.FixturesByDateResp, error) {
	body, err := c.get(
		fmt.Sprintf(
			fixturesByDateEndpoint,
			leagueID,
			season,
			c.formatDate(fromDate),
			c.formatDate(toDate),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting fixtures by date and league: %w", err)
	}

	var resp models.FixturesByDateResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling fixtures response by date and league: %w",
			err,
		)
	}
	return &resp, nil
}

// FixtureStatistics returns statistics for a given fixture
/*
	- fixture (Type: integer)(Required)
	The ID of the fixture. Value format: 43.
	- team (Type: integer)
	The ID of the team. Value format: 85.
*/
func (c *Client) FixtureStatistics(
	params map[string]interface{},
) (*models.FixturesStatisticsResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, fixtureStatisticsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting fixtures statistics: %w", err)
	}

	var resp models.FixturesStatisticsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling fixture statistics response: %w",
			err,
		)
	}
	return &resp, nil
}

// FixturesPlayer returns fixture's player data
/*
	- fixture (Type: integer)(Required)
	The ID of the fixture. Value format: 43.
	- team (Type: integer)
	The ID of the team. Value format: 85.
*/
func (c *Client) FixturesPlayer(
	params map[string]interface{},
) (*models.FixturesPlayersResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, fixturePlayerEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting fixtures statistics: %w", err)
	}

	var resp models.FixturesPlayersResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling fixtures player response: %w",
			err,
		)
	}
	return &resp, nil
}
