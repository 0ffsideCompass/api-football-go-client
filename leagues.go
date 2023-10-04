package client

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	leaguesEndpoint       = "leagues"
	leagueSeasonsEndpoint = "leagueSeasons"
)

// Leagues returns all leagues
/*
	- id (Type: integer)
	  The ID of the league. Value format: 2
	- name (Type: string)
	  The name of the league. Value format: Premier League
	-country (Type: string)
	  The name of the country. Value format: England
	- season (Type: integer) (should be 4 digits)
	  The year of the season. Value format: 2020
	- team (Type: integer)
	  The ID of the team. Value format: 85
	- type (Type: string)
	  The type of the league. Value format: League
	- current (Type: string) Enum: (true, false)
	  Return the list of active seasons or the last one of each competition. Value format: "true"
	- search (Type: string) (should be 3 characters or more)
	  The name of the league. Value format: Premier League
	- last (Type: integer) (<= 2 digis)
	  The X last leagues/cups added in the API
	- code (Type: string) (2 characters)
	  The code of the country. Value format: FR
*/
func (c *Client) Leagues(
	params map[string]interface{},
) (*models.LeaguesResponse, error) {
	// Validate the parameters
	if err := validateLeaguesParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, leaguesEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting leagues response: %w", err)
	}

	var resp models.LeaguesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling leagues response: %w",
			err,
		)
	}
	return &resp, nil
}

func validateLeaguesParams(params map[string]interface{}) error {
	// Validate 'id' parameter
	if id, ok := params["id"].(float64); ok {
		if id != float64(int(id)) {
			return fmt.Errorf("'id' must be an integer")
		}
	}

	// Validate 'name' parameter
	if name, ok := params["name"].(string); ok {
		if len(name) == 0 {
			return fmt.Errorf("'name' must not be empty")
		}
	}

	// Validate 'country' parameter
	if country, ok := params["country"].(string); ok {
		if len(country) == 0 {
			return fmt.Errorf("'country' must not be empty")
		}
	}

	// Validate 'season' parameter
	if season, ok := params["season"].(float64); ok {
		seasonStr := strconv.Itoa(int(season))
		if len(seasonStr) != 4 {
			return fmt.Errorf("'season' must be 4 digits")
		}
	}

	// Validate 'team' parameter
	if teamID, ok := params["team"].(float64); ok {
		if teamID != float64(int(teamID)) {
			return fmt.Errorf("'team' must be an integer")
		}
	}

	// Validate 'type' parameter
	if typeStr, ok := params["type"].(string); ok {
		if typeStr != "League" {
			return fmt.Errorf("'type' must be 'League'")
		}
	}

	// Validate 'current' parameter
	if current, ok := params["current"].(string); ok {
		if current != "true" && current != "false" {
			return fmt.Errorf("'current' must be 'true' or 'false'")
		}
	}

	// Validate 'search' parameter
	if search, ok := params["search"].(string); ok {
		if len(search) < 3 {
			return fmt.Errorf("'search' must be at least 3 characters long")
		}
	}

	// Validate 'last' parameter
	if last, ok := params["last"].(float64); ok {
		if last != float64(int(last)) || last >= 100 {
			return fmt.Errorf("'last' must be an integer with at most 2 digits")
		}
	}

	// Validate 'code' parameter
	if code, ok := params["code"].(string); ok {
		if len(code) != 2 || !isAlpha(code) {
			return fmt.Errorf("'code' must be a 2-character alphabetic string")
		}
	}

	return nil
}

// Helper function to check if a string contains only alphabetic characters
func isAlpha(s string) bool {
	alphaRegex := regexp.MustCompile("^[a-zA-Z]+$")
	return alphaRegex.MatchString(s)
}
