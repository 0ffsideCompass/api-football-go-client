package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	teamsEndpoint           = "teams"
	teamsStatisticsEndpoint = "teams/statistics"
)

// Teams hits the /teams endpoint
/*
	-id: (Type: integer)
	  The ID of the team. Value format: 85
	- name: (Type: string)
	  The name of the team. Value format: Arsenal
	- league: (Type: integer)
	  The ID of the league. Value format: 2
	- season: (Type: integer) (should be 4 digits)
	  The year of the season. Value format: 2020
	- country: (Type: string)
	  The name of the country. Value format: England
	- code: (Type: string) (3 characters)
	  The code of the team. Value format: ARS
	- venue: (Type: integer)
	  The ID of the venue. Value format: 18
	- search: (Type: string) (should be 3 characters or more)
	  The name of the team. Value format: Arsenal

	At least one of the following parameters must be passed
*/
func (c *Client) Teams(
	params map[string]interface{},
) (*models.TeamsResponse, error) {
	// Validate the parameters
	if err := validateTeamsParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, teamsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting teams: %w", err)
	}

	var resp models.TeamsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling teams response: %w",
			err,
		)
	}

	return &resp, nil
}

// TeamsStatistics hits the /teams/statistics endpoint
/*
	- league (Type: integer) (required)
	  The ID of the league. Value format: 2
	- season (Type: integer) (required) (should be 4 digits)
	  The year of the season. Value format: 2020
	- team (Type: integer) (required)
	  The ID of the team. Value format: 85
	- date (Type: string) (format: YYYY-MM-DD)
	  The limit date. Value format: 2020-12-01
*/
func (c *Client) TeamsStatistics(
	params map[string]interface{},
) (*models.TeamsStatisticsResponse, error) {
	endpointURL := fmt.Sprintf("%s%s", c.Domain, teamsStatisticsEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting teams statistics: %w", err)
	}

	var resp models.TeamsStatisticsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling teams statistics response: %w",
			err,
		)
	}

	return &resp, nil
}

func validateTeamsParams(params map[string]interface{}) error {
	// Define an array of parameter names to check for their presence
	requiredParams := []string{"id", "name", "league", "season", "country", "code", "venue", "search"}

	// Iterate through the required parameters and check if at least one is provided
	found := false
	for _, paramName := range requiredParams {
		if _, hasParam := params[paramName]; hasParam {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("at least one of 'id', 'name', 'league', 'season', 'country', 'code', 'venue', or 'search' must be provided")
	}

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

	// Validate 'league' parameter
	if leagueID, ok := params["league"].(float64); ok {
		if leagueID != float64(int(leagueID)) {
			return fmt.Errorf("'league' must be an integer")
		}
	}

	// Validate 'season' parameter
	if season, ok := params["season"].(float64); ok {
		seasonStr := strconv.Itoa(int(season))
		if len(seasonStr) != 4 {
			return fmt.Errorf("'season' must be 4 digits")
		}
	}

	// Validate 'country' parameter
	if country, ok := params["country"].(string); ok {
		if len(country) == 0 {
			return fmt.Errorf("'country' must not be empty")
		}
	}

	// Validate 'code' parameter
	if code, ok := params["code"].(string); ok {
		if len(code) != 3 {
			return fmt.Errorf("'code' must be 3 characters long")
		}
	}

	// Validate 'venue' parameter
	if venueID, ok := params["venue"].(float64); ok {
		if venueID != float64(int(venueID)) {
			return fmt.Errorf("'venue' must be an integer")
		}
	}

	// Validate 'search' parameter
	if search, ok := params["search"].(string); ok {
		if len(search) < 3 {
			return fmt.Errorf("'search' must be at least 3 characters long")
		}
	}

	return nil
}

// validateTeamsStatisticsParams validates the parameters for the TeamsStatistics endpoint.
func validateTeamsStatisticsParams(params map[string]interface{}) error {
	// Validate 'league' parameter (required)
	if leagueID, ok := params["league"].(float64); ok {
		if leagueID != float64(int(leagueID)) {
			return fmt.Errorf("'league' must be an integer")
		}
	} else {
		return fmt.Errorf("'league' is required")
	}

	// Validate 'season' parameter (required)
	if season, ok := params["season"].(float64); ok {
		seasonStr := strconv.Itoa(int(season))
		if len(seasonStr) != 4 {
			return fmt.Errorf("'season' must be 4 digits")
		}
	} else {
		return fmt.Errorf("'season' is required")
	}

	// Validate 'team' parameter (required)
	if teamID, ok := params["team"].(float64); ok {
		if teamID != float64(int(teamID)) {
			return fmt.Errorf("'team' must be an integer")
		}
	} else {
		return fmt.Errorf("'team' is required")
	}

	// Validate 'date' parameter
	if dateStr, ok := params["date"].(string); ok {
		// Check the date format (YYYY-MM-DD)
		if !isValidDateFormat(dateStr) {
			return fmt.Errorf("'date' must be in the format 'YYYY-MM-DD'")
		}
	}

	return nil
}
