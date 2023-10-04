package client

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	injuriesEndpoint = "injuries"
)

// Injuries hits the /injuries endpoint
/*
	- league: (Type: integer)
	  The ID of the league. Value format: 2
	- season: (Type: integer) (should be 4 digits)
	  The year of the season. Value format: 2020
	- fixture: (Type: integer)
	  The ID of the fixture. Value format: 157304
	- team: (Type: integer)
	  The ID of the team. Value format: 85
	- player: (Type: integer)
	  The ID of the player. Value format: 85
	- date: (Type: string) (format: YYYY-MM-DD)
	  The date of the fixture. Value format: 2020-12-01

	if league is provided, season is required
	if team is provided, season is required
	if player is provided, season is required
*/
func (c *Client) Injuries(
	params map[string]interface{},
) (*models.InjuriesResponse, error) {
	// Validate the parameters
	if err := validateInjuriesParams(params); err != nil {
		return nil, err
	}

	endpointURL := fmt.Sprintf("%s%s", c.Domain, injuriesEndpoint)
	body, err := c.get(
		c.buildURL(
			endpointURL,
			params,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting injuries: %w", err)
	}

	var resp models.InjuriesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling injuries response: %w",
			err,
		)
	}

	return &resp, nil
}

func validateInjuriesParams(params map[string]interface{}) error {
	// Validate 'league' parameter
	if leagueID, ok := params["league"].(float64); ok {
		if leagueID != float64(int(leagueID)) {
			return fmt.Errorf("'league' must be an integer")
		}

		// If 'league' is provided, 'season' is required
		if _, hasSeason := params["season"]; !hasSeason {
			return fmt.Errorf("'season' is required when 'league' is provided")
		}
	}

	// Validate 'season' parameter
	if season, ok := params["season"].(float64); ok {
		seasonStr := strconv.Itoa(int(season))
		if len(seasonStr) != 4 {
			return fmt.Errorf("'season' must be 4 digits")
		}
	}

	// Validate 'fixture' parameter
	if fixtureID, ok := params["fixture"].(float64); ok {
		if fixtureID != float64(int(fixtureID)) {
			return fmt.Errorf("'fixture' must be an integer")
		}
	}

	// Validate 'team' parameter
	if teamID, ok := params["team"].(float64); ok {
		if teamID != float64(int(teamID)) {
			return fmt.Errorf("'team' must be an integer")
		}

		// If 'team' is provided, 'season' is required
		if _, hasSeason := params["season"]; !hasSeason {
			return fmt.Errorf("'season' is required when 'team' is provided")
		}
	}

	// Validate 'player' parameter
	if playerID, ok := params["player"].(float64); ok {
		if playerID != float64(int(playerID)) {
			return fmt.Errorf("'player' must be an integer")
		}

		// If 'player' is provided, 'season' is required
		if _, hasSeason := params["season"]; !hasSeason {
			return fmt.Errorf("'season' is required when 'player' is provided")
		}
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

// Helper function to check the date format (YYYY-MM-DD)
func isValidDateFormat(dateStr string) bool {
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return dateRegex.MatchString(dateStr)
}
