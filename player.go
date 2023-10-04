package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	playersEndpoint               = "players"
	playersSeasonsEndpoint        = "players/seasons"
	playersSquadsEndpoint         = "players/squads"
	playersTopScorersEndpoint     = "players/topscorers"
	playersTopAssistsEndpoint     = "players/topassists"
	playersTopYellowCardsEndpoint = "players/topyellowcards"
	playersTopRedCardsEndpoint    = "players/topredcards"
)

// PlayersSeasons returns the seasons for a given player
/*
	- player (Type: integer)
	  The ID of the player. Value format: 85
*/
func (c *Client) PlayersSeasons(
	params map[string]interface{},
) (*models.PlayersSeasonsResponse, error) {

	// Validate the parameters
	if _, hasPlayer := params["player"]; !hasPlayer {
		return nil, fmt.Errorf("at least one of 'coach' or 'player' must be provided")
	}

	val, ok := params["player"].(int)
	if !ok {
		return nil, fmt.Errorf("player mustexist")
	}
	if val != int(int(val)) {
		return nil, fmt.Errorf("%d must be an integer", val)
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, playersSeasonsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting players seasons: %w", err)
	}

	var resp models.PlayersSeasonsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling players seasons response: %w",
			err,
		)
	}

	return &resp, nil
}

// Players returns the players data
/*
	- id (Type: integer)
	The ID of the player. Value format: 85
	- team (Type: integer)
	The ID of the team. Value format: 85
	- season (Type: integer) (4 characters) Requires one of the fields Id, League, or Team to be present.
	The year of the season. Value format: 2020
	- search (Type: string) (>= 4 characters) Requires: One of the fields League or Team to be present.
	The name of the player. Value format: Ronaldo
	- page (Type: integer)
	Use for pagination.
*/

func (c *Client) Players(
	params map[string]interface{},
) (*models.PlayersResponse, error) {

	// Validate the parameters
	if err := validatePlayersParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, playersEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting players response: %w", err)
	}

	var resp models.PlayersResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling players response: %w",
			err,
		)
	}

	return &resp, nil
}

// PlayersSquads returns the players data for a given team
/*
	- team (Type: integer)
	  The ID of the team. Value format: 85
	- player (Type: integer)
	  The ID of the player. Value format: 85
*/

func (c *Client) PlayersSquads(
	params map[string]interface{},
) (*models.PlayersSquadsResponse, error) {

	val, ok := params["team"].(int)
	if !ok {
		return nil, fmt.Errorf("team must exist")
	} else {
		if val != int(int(val)) {
			return nil, fmt.Errorf("%d must be an integer", val)
		}
	}

	val, ok = params["player"].(int)
	if !ok {
		return nil, fmt.Errorf("player must exist")
	} else {
		if val != int(int(val)) {
			return nil, fmt.Errorf("%d must be an integer", val)
		}
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, playersSquadsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting players squads: %w", err)
	}

	var resp models.PlayersSquadsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling players squads response: %w",
			err,
		)
	}

	return &resp, nil
}

// PlayersTopScorers returns the top scorers for a given league and season
/*
	- league (Type: integer)(Required)
	  The ID of the league. Value format: 2
	- season (Type: integer)(Required) (4 digits)
	  The year of the season. Value format: 2020
*/
func (c *Client) PlayersTopScorers(
	params map[string]interface{},
) (*models.PlayersTopResponse, error) {

	// Validate the parameters
	if err := validateTopParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, playersTopScorersEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting players topscorers: %w", err)
	}

	var resp models.PlayersTopResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling players topscorers response: %w",
			err,
		)
	}

	return &resp, nil
}

// PlayersTopAssists returns the top assists for a given league and season
/*
	- league (Type: integer)(Required)
	  The ID of the league. Value format: 2
	- season (Type: integer)(Required) (4 digits)
	  The year of the season. Value format: 2020
*/
func (c *Client) PlayersTopAssists(
	params map[string]interface{},
) (*models.PlayersTopResponse, error) {

	// Validate the parameters
	if err := validateTopParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, playersTopAssistsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting players topscorers: %w", err)
	}

	var resp models.PlayersTopResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling players topscorers response: %w",
			err,
		)
	}

	return &resp, nil
}

// PlayersTopYellowCards returns the top assists for a given league and season
/*
	- league (Type: integer)(Required)
	  The ID of the league. Value format: 2
	- season (Type: integer)(Required) (4 digits)
	  The year of the season. Value format: 2020
*/
func (c *Client) PlayersTopYellowCards(
	params map[string]interface{},
) (*models.PlayersTopResponse, error) {

	// Validate the parameters
	if err := validateTopParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, playersTopYellowCardsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting players topscorers: %w", err)
	}

	var resp models.PlayersTopResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling players topscorers response: %w",
			err,
		)
	}

	return &resp, nil
}

// PlayersTopRedCards returns the top assists for a given league and season
/*
	- league (Type: integer)(Required)
	  The ID of the league. Value format: 2
	- season (Type: integer)(Required) (4 digits)
	  The year of the season. Value format: 2020
*/
func (c *Client) PlayersTopRedCards(
	params map[string]interface{},
) (*models.PlayersTopResponse, error) {

	// Validate the parameters
	if err := validateTopParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, playersTopRedCardsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting players topscorers: %w", err)
	}

	var resp models.PlayersTopResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling players topscorers response: %w",
			err,
		)
	}

	return &resp, nil
}

func validatePlayersParams(params map[string]interface{}) error {
	// Check 'id'
	id, idExists := params["id"].(int)
	if idExists && id < 1 {
		return errors.New("invalid player id parameter")
	}

	// Check 'team'
	team, teamExists := params["team"].(int)
	if teamExists && team < 1 {
		return errors.New("invalid team parameter")
	}

	// Check 'season'
	if season, seasonExists := params["season"].(int); seasonExists {
		if len(fmt.Sprintf("%d", season)) != 4 {
			return errors.New("season should be 4 characters long")
		}
		if !idExists && !teamExists {
			_, leagueExists := params["league"].(int)
			if !leagueExists {
				return errors.New("for 'season', one of the fields Id, League, or Team should be present")
			}
		}
	}

	// Check 'search'
	if search, searchExists := params["search"].(string); searchExists {
		if utf8.RuneCountInString(search) < 4 {
			return errors.New("search parameter should be 4 characters or more")
		}
		if !teamExists {
			_, leagueExists := params["league"].(int)
			if !leagueExists {
				return errors.New("for 'search', one of the fields League or Team should be present")
			}
		}
	}

	// Check 'page' (optional, just ensure it's a positive integer if exists)
	if page, pageExists := params["page"].(int); pageExists && page < 1 {
		return errors.New("invalid page parameter")
	}

	return nil
}

func validateTopParams(params map[string]interface{}) error {
	// Check 'league'
	league, leagueExists := params["league"].(int)
	if !leagueExists || league < 1 {
		return errors.New("league parameter is required and should be a positive integer")
	}

	// Check 'season'
	if season, seasonExists := params["season"].(int); seasonExists {
		if len(fmt.Sprintf("%d", season)) != 4 {
			return errors.New("season should be 4 characters long")
		}
	} else {
		return errors.New("season parameter is required")
	}

	return nil
}
