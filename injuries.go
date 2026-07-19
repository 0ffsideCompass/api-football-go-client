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

	// seasonDigits is the number of digits the API expects for a season (YYYY).
	seasonDigits = 4
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
	params map[string]any,
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

func validateInjuriesParams(params map[string]any) error {
	// Validate integer parameters. Values may arrive as int (Go callers) or
	// float64 (values decoded from JSON).
	intParams := []struct {
		key         string
		needsSeason bool
	}{
		{key: "league", needsSeason: true},
		{key: "fixture", needsSeason: false},
		{key: "team", needsSeason: true},
		{key: "player", needsSeason: true},
	}
	for _, p := range intParams {
		value, ok := params[p.key]
		if !ok {
			continue
		}
		if _, err := asInt(value); err != nil {
			return fmt.Errorf("'%s' must be an integer", p.key)
		}
		if p.needsSeason {
			if _, hasSeason := params["season"]; !hasSeason {
				return fmt.Errorf("'season' is required when '%s' is provided", p.key)
			}
		}
	}

	// Validate 'season' parameter
	if seasonVal, ok := params["season"]; ok {
		season, err := asInt(seasonVal)
		if err != nil {
			return fmt.Errorf("'season' must be an integer")
		}
		if len(strconv.Itoa(season)) != seasonDigits {
			return fmt.Errorf("'season' must be %d digits", seasonDigits)
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

// asInt converts an int, a whole-number float64, or a numeric string to int.
// All three shapes occur in practice: Go callers pass ints, JSON-decoded
// params arrive as float64, and callers forwarding path/query values pass
// strings.
func asInt(value any) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case float64:
		if v != float64(int(v)) {
			return 0, fmt.Errorf("not an integer")
		}
		return int(v), nil
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("not an integer")
		}
		return n, nil
	default:
		return 0, fmt.Errorf("not an integer")
	}
}

// Helper function to check the date format (YYYY-MM-DD)
func isValidDateFormat(dateStr string) bool {
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return dateRegex.MatchString(dateStr)
}
