package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

// endpointCalls invokes every endpoint method with minimal valid parameters,
// so shared error-path behavior can be asserted across all of them.
var endpointCalls = []struct {
	name    string
	rawBody bool // returns the raw body without JSON unmarshalling
	call    func(c *client.Client) error
}{
	{name: "Timezone", call: func(c *client.Client) error { _, err := c.Timezone(); return err }},
	{name: "Countries", call: func(c *client.Client) error { _, err := c.Countries(nil); return err }},
	{name: "Leagues", call: func(c *client.Client) error { _, err := c.Leagues(nil); return err }},
	{name: "LeaguesSeasons", call: func(c *client.Client) error { _, err := c.LeaguesSeasons(); return err }},
	{name: "Teams", call: func(c *client.Client) error { _, err := c.Teams(nil); return err }},
	{name: "TeamsStatistics", call: func(c *client.Client) error { _, err := c.TeamsStatistics(nil); return err }},
	{name: "TeamsSeasons", call: func(c *client.Client) error {
		_, err := c.TeamsSeasons(map[string]any{"team": 33})
		return err
	}},
	{name: "TeamsCountries", call: func(c *client.Client) error { _, err := c.TeamsCountries(); return err }},
	{name: "Venues", call: func(c *client.Client) error { _, err := c.Venues(nil); return err }},
	{name: "Standings", call: func(c *client.Client) error { _, err := c.Standings(nil); return err }},
	{name: "Fixture", call: func(c *client.Client) error { _, err := c.Fixture(nil); return err }},
	{name: "FixturesRounds", call: func(c *client.Client) error {
		_, err := c.FixturesRounds(map[string]any{"league": 39, "season": 2023})
		return err
	}},
	{name: "FixturesLineups", call: func(c *client.Client) error { _, err := c.FixturesLineups(nil); return err }},
	{name: "FixturesEvents", call: func(c *client.Client) error { _, err := c.FixturesEvents(nil); return err }},
	{name: "FixtureHeadToHead", call: func(c *client.Client) error { _, err := c.FixtureHeadToHead(nil); return err }},
	{name: "FixtureByDateAndLeague", call: func(c *client.Client) error {
		day := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
		_, err := c.FixtureByDateAndLeague(39, 2023, day, day)
		return err
	}},
	{name: "FixtureStatistics", call: func(c *client.Client) error { _, err := c.FixtureStatistics(nil); return err }},
	{name: "FixturesPlayer", call: func(c *client.Client) error { _, err := c.FixturesPlayer(nil); return err }},
	{name: "Injuries", call: func(c *client.Client) error {
		_, err := c.Injuries(map[string]any{"fixture": 157304})
		return err
	}},
	{name: "Predictions", call: func(c *client.Client) error {
		_, err := c.Predictions(map[string]any{"fixture": 198772})
		return err
	}},
	{name: "Coachs", call: func(c *client.Client) error {
		_, err := c.Coachs(map[string]any{"id": 1})
		return err
	}},
	{name: "PlayersSeasons", call: func(c *client.Client) error {
		_, err := c.PlayersSeasons(map[string]any{"player": 276})
		return err
	}},
	{name: "Players", call: func(c *client.Client) error { _, err := c.Players(nil); return err }},
	{name: "PlayersProfiles", call: func(c *client.Client) error { _, err := c.PlayersProfiles(nil); return err }},
	{name: "PlayersSquads", call: func(c *client.Client) error {
		_, err := c.PlayersSquads(map[string]any{"team": 33})
		return err
	}},
	{name: "PlayersTeams", call: func(c *client.Client) error {
		_, err := c.PlayersTeams(map[string]any{"player": 276})
		return err
	}},
	{name: "PlayersTopScorers", call: func(c *client.Client) error { _, err := c.PlayersTopScorers(nil); return err }},
	{name: "PlayersTopAssists", call: func(c *client.Client) error { _, err := c.PlayersTopAssists(nil); return err }},
	{name: "PlayersTopYellowCards", call: func(c *client.Client) error { _, err := c.PlayersTopYellowCards(nil); return err }},
	{name: "PlayersTopRedCards", call: func(c *client.Client) error { _, err := c.PlayersTopRedCards(nil); return err }},
	{name: "Transfers", call: func(c *client.Client) error {
		_, err := c.Transfers(map[string]any{"player": 276})
		return err
	}},
	{name: "Trophies", call: func(c *client.Client) error {
		_, err := c.Trophies(map[string]any{"player": 276})
		return err
	}},
	{name: "Sidelined", call: func(c *client.Client) error {
		_, err := c.Sidelined(map[string]any{"player": 276})
		return err
	}},
	{name: "Odds", call: func(c *client.Client) error { _, err := c.Odds(nil); return err }},
	{name: "OddsMapping", call: func(c *client.Client) error { _, err := c.OddsMapping(nil); return err }},
	{name: "OddsBookmakers", call: func(c *client.Client) error { _, err := c.OddsBookmakers(nil); return err }},
	{name: "OddsBets", call: func(c *client.Client) error { _, err := c.OddsBets(nil); return err }},
	{name: "OddsLive", call: func(c *client.Client) error { _, err := c.OddsLive(nil); return err }},
	{name: "OddsLiveBets", call: func(c *client.Client) error { _, err := c.OddsLiveBets(nil); return err }},
	{name: "Search", rawBody: true, call: func(c *client.Client) error {
		_, err := c.Search("Arsenal", client.Team)
		return err
	}},
}

func TestEndpointHTTPErrors(t *testing.T) {
	for _, tt := range endpointCalls {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(bytes.NewBufferString(`{"message": "server error"}`)),
				},
			}
			apiClient, err := client.New("test-api-key", mock)
			assert.NoError(t, err)

			err = tt.call(apiClient)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "API request failed with status 500")
		})
	}
}

func TestEndpointInvalidJSON(t *testing.T) {
	for _, tt := range endpointCalls {
		if tt.rawBody {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			apiClient := newTestClient(t, `{"invalid": json}`)

			err := tt.call(apiClient)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "error unmarshalling")
		})
	}
}

func TestCoachsValidation(t *testing.T) {
	apiClient, err := client.New("test-api-key", newMockClient())
	assert.NoError(t, err)

	_, err = apiClient.Coachs(map[string]any{})
	assert.EqualError(t, err, "at least one of 'team', 'id', or 'search' must be provided")

	_, err = apiClient.Coachs(map[string]any{"search": "ab"})
	assert.EqualError(t, err, "'search' must be at least 3 characters long")

	_, err = apiClient.Coachs(map[string]any{"search": "abc"})
	assert.NoError(t, err)
}

func TestInjuriesValidation(t *testing.T) {
	tests := []struct {
		name    string
		params  map[string]any
		wantErr string
	}{
		{
			name:    "valid float64 params",
			params:  map[string]any{"league": 39.0, "season": 2023.0},
			wantErr: "",
		},
		{
			name:    "valid numeric string params",
			params:  map[string]any{"fixture": "1581037"},
			wantErr: "",
		},
		{
			name:    "non-numeric string fixture",
			params:  map[string]any{"fixture": "abc"},
			wantErr: "'fixture' must be an integer",
		},
		{
			name:    "non-integer float64 league",
			params:  map[string]any{"league": 39.5, "season": 2023},
			wantErr: "'league' must be an integer",
		},
		{
			name:    "non-integer season",
			params:  map[string]any{"season": 2023.5},
			wantErr: "'season' must be an integer",
		},
		{
			name:    "season with wrong number of digits",
			params:  map[string]any{"season": 999},
			wantErr: "'season' must be 4 digits",
		},
		{
			name:    "valid date format",
			params:  map[string]any{"date": "2023-01-01"},
			wantErr: "",
		},
		{
			name:    "invalid date format",
			params:  map[string]any{"date": "01-01-2023 extra"},
			wantErr: "'date' must be in the format 'YYYY-MM-DD'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClient := newTestClient(t, `{"response": [], "errors": [], "results": 0}`)

			_, err := apiClient.Injuries(tt.params)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
