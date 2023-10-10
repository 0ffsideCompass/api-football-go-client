package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

	// Validate the parameters
	if err := validateFixturesEventsLineupsParams(params); err != nil {
		return nil, err
	}

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

	// Validate the parameters
	if err := validateFixturesEventsLineupsParams(params); err != nil {
		return nil, err
	}

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

	// Validate the parameters
	if err := validateFixtureParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s/%s", c.Domain, fixtureEndpoint)

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
		return nil, err
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

	// Validate the parameters
	if err := validateFixtureStatisticsPlayerParams(params); err != nil {
		return nil, err
	}

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
			"error unmarshalling fixtures statistics response: %w",
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
	// Validate the parameters
	if err := validateFixtureStatisticsPlayerParams(params); err != nil {
		return nil, err
	}

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

func validateFixtureParams(params map[string]interface{}) error {

	// Validate ID
	if id, exists := params["id"].(int); exists && id < 1 {
		return errors.New("invalid id parameter")
	}

	// Validate IDs
	if ids, exists := params["ids"].(string); exists {
		idList := strings.Split(ids, "-")
		if len(idList) > 20 {
			return errors.New("maximum of 20 fixture IDs allowed")
		}
		for _, id := range idList {
			if _, err := strconv.Atoi(id); err != nil {
				return errors.New("invalid id in ids parameter")
			}
		}
	}

	// Validate date
	if date, exists := params["date"].(string); exists {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			return errors.New("invalid date format")
		}
	}

	// Validate from and to
	if fromDate, exists := params["from"].(string); exists {
		_, err := time.Parse("2006-01-02", fromDate)
		if err != nil {
			return errors.New("invalid from date format")
		}
	}
	if toDate, exists := params["to"].(string); exists {
		_, err := time.Parse("2006-01-02", toDate)
		if err != nil {
			return errors.New("invalid to date format")
		}
	}

	// Validate season
	if season, exists := params["season"].(int); exists && (season < 1000 || season > 9999) {
		return errors.New("invalid season format. Season should be 4 digits")
	}

	// Validate last and next
	if last, exists := params["last"].(int); exists && (last < 0 || last > 99) {
		return errors.New("invalid last parameter. Value should be an integer with <= 2 characters")
	}
	if next, exists := params["next"].(int); exists && (next < 0 || next > 99) {
		return errors.New("invalid next parameter. Value should be an integer with <= 2 characters")
	}

	// Validate status
	if status, exists := params["status"].(string); exists {
		statusList := strings.Split(status, "-")
		validStatuses := map[string]bool{
			"NS":  true,
			"PST": true,
			"FT":  true,
		}
		for _, s := range statusList {
			if _, valid := validStatuses[s]; !valid {
				return errors.New("invalid status value")
			}
		}
	}

	return nil
}

func validateFixtureHeadToHeadParams(params map[string]interface{}) error {

	// Validate h2h (head-to-head)
	if h2h, exists := params["h2h"].(string); exists {
		teamIds := strings.Split(h2h, "-")
		if len(teamIds) != 2 {
			return errors.New("h2h parameter must have exactly two team IDs separated by '-'")
		}
		for _, id := range teamIds {
			if _, err := strconv.Atoi(id); err != nil {
				return errors.New("invalid team ID in h2h parameter")
			}
		}
	} else {
		return errors.New("h2h parameter is required")
	}

	// Validate date
	if date, exists := params["date"].(string); exists {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			return errors.New("invalid date format")
		}
	}

	// Validate from and to
	if fromDate, exists := params["from"].(string); exists {
		_, err := time.Parse("2006-01-02", fromDate)
		if err != nil {
			return errors.New("invalid from date format")
		}
	}
	if toDate, exists := params["to"].(string); exists {
		_, err := time.Parse("2006-01-02", toDate)
		if err != nil {
			return errors.New("invalid to date format")
		}
	}

	// Validate season
	if season, exists := params["season"].(int); exists && (season < 1000 || season > 9999) {
		return errors.New("invalid season format. Season should be 4 digits")
	}

	// Validate status
	if status, exists := params["status"].(string); exists {
		statusList := strings.Split(status, "-")
		validStatuses := map[string]bool{
			"NS":  true,
			"PST": true,
			"FT":  true,
		}
		for _, s := range statusList {
			if _, valid := validStatuses[s]; !valid {
				return errors.New("invalid status value")
			}
		}
	}

	return nil
}

func validateFixtureStatisticsPlayerParams(params map[string]interface{}) error {
	// Validate fixture
	if fixture, exists := params["fixture"].(int); !exists || fixture < 1 {
		return errors.New("invalid or missing fixture parameter")
	}

	// Validate team (optional, so just validate if exists)
	if team, exists := params["team"].(int); exists && team < 1 {
		return errors.New("invalid team parameter")
	}

	return nil
}

func validateFixturesEventsLineupsParams(params map[string]interface{}) error {
	// Validate fixture
	if fixture, exists := params["fixture"].(int); !exists || fixture < 1 {
		return errors.New("invalid or missing fixture parameter")
	}

	// Validate team (optional, so just validate if exists)
	if team, exists := params["team"].(int); exists && team < 1 {
		return errors.New("invalid team parameter")
	}

	// Validate player (optional, so just validate if exists)
	if player, exists := params["player"].(int); exists && player < 1 {
		return errors.New("invalid player parameter")
	}

	return nil
}
