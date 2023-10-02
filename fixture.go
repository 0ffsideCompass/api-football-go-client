package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	fixturesByDateEndpoint    = "fixtures?league=%d&season=%d&from=%s&to=%s"
	fixtureHeadToHeadEndpoint = "fixtures/headtohead?h2h=%d-%d"
	fixtureRounds             = "fixtures/rounds?league=%d&season=%d" //TODO not important yet
	fixture                   = "fixtures"
)

// Fixture returns all fixtures for parameters passed
/*
============================== QUERY PARAMETERS ==============================

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

- season: (Type: integer)
  The season of the league. Format: 4 characters YYYY.

- team: (Type: integer)
  The ID of the team.

- last: (Type: integer)
  For the X last fixtures. Value should be an integer with <= 2 characters.

- next: (Type: integer)
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
) (*models.FixturesByDateResp, error) {
	endpointURL := fmt.Sprintf("%s/%s", c.Domain, fixture)

	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting fixtures: %w", err)
	}
	fmt.Println(string(body))
	var resp models.FixturesByDateResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling fixtures response: %w",
			err,
		)
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

func (c *Client) FixtureHeadToHead(
	teamOneID,
	teamTwoID int,
) (*models.FixtureHeadToHeadResp, error) {
	body, err := c.get(
		fmt.Sprintf(
			fixtureHeadToHeadEndpoint,
			teamOneID,
			teamTwoID,
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
