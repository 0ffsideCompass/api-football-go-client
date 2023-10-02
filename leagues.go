package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	leaguesEndpoint = "leagues"
)

// Leagues returns all leagues

func (c *Client) Leagues() (*models.LeaguesResponse, error) {
	body, err := c.get(leaguesEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error getting leagues: %w", err)
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
