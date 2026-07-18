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

func newMockClient() *MockHTTPClient {
	return &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(
				`{"response": [], "errors": [], "results": 0, "paging": {"current": 1, "total": 1}}`,
			)),
		},
	}
}

// TestRequestURLs verifies that each method hits the endpoint it is named after.
func TestRequestURLs(t *testing.T) {
	tests := []struct {
		name        string
		call        func(c *client.Client) error
		expectedURL string
	}{
		{
			name: "FixturesLineups hits fixtures/lineups",
			call: func(c *client.Client) error {
				_, err := c.FixturesLineups(map[string]any{"fixture": 43})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/fixtures/lineups?fixture=43",
		},
		{
			name: "FixturesEvents hits fixtures/events",
			call: func(c *client.Client) error {
				_, err := c.FixturesEvents(map[string]any{"fixture": 43})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/fixtures/events?fixture=43",
		},
		{
			name: "Transfers hits transfers",
			call: func(c *client.Client) error {
				_, err := c.Transfers(map[string]any{"player": 276})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/transfers?player=276",
		},
		{
			name: "Trophies hits trophies",
			call: func(c *client.Client) error {
				_, err := c.Trophies(map[string]any{"player": 276})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/trophies?player=276",
		},
		{
			name: "FixtureByDateAndLeague includes the domain",
			call: func(c *client.Client) error {
				from := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
				to := time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC)
				_, err := c.FixtureByDateAndLeague(39, 2023, from, to)
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/fixtures?league=39&season=2023&from=2023-08-01&to=2023-08-31",
		},
		{
			name: "Search teams uses the search query parameter",
			call: func(c *client.Client) error {
				_, err := c.Search("Manchester United", client.Team)
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/teams?search=Manchester+United",
		},
		{
			name: "Search leagues uses the search query parameter",
			call: func(c *client.Client) error {
				_, err := c.Search("Premier", client.League)
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/leagues?search=Premier",
		},
		{
			name: "Search players uses the search query parameter",
			call: func(c *client.Client) error {
				_, err := c.Search("Ronaldo", client.Player)
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/players?search=Ronaldo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := newMockClient()
			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			assert.NoError(t, tt.call(apiClient))
			assert.NotNil(t, mockClient.LastRequest)
			assert.Equal(t, tt.expectedURL, mockClient.LastRequest.URL.String())
		})
	}
}

func TestRequestHeaders(t *testing.T) {
	mockClient := newMockClient()
	apiClient, err := client.New("test-api-key", mockClient)
	assert.NoError(t, err)

	_, err = apiClient.Leagues(map[string]any{"season": 2023})
	assert.NoError(t, err)

	headers := mockClient.LastRequest.Header
	assert.Equal(t, "test-api-key", headers.Get("X-RapidAPI-Key"))
	assert.Equal(t, "api-football-v1.p.rapidapi.com", headers.Get("X-RapidAPI-Host"))
	assert.Equal(t, "test-api-key", headers.Get("x-apisports-key"))
}

func TestNewWithDomain(t *testing.T) {
	t.Run("appends a trailing slash to the domain", func(t *testing.T) {
		apiClient, err := client.NewWithDomain("key", "https://example.com/v3", &http.Client{})
		assert.NoError(t, err)
		assert.Equal(t, "https://example.com/v3/", apiClient.Domain)
	})

	t.Run("keeps an existing trailing slash", func(t *testing.T) {
		apiClient, err := client.NewWithDomain("key", "https://example.com/v3/", &http.Client{})
		assert.NoError(t, err)
		assert.Equal(t, "https://example.com/v3/", apiClient.Domain)
	})

	t.Run("rejects an empty domain", func(t *testing.T) {
		_, err := client.NewWithDomain("key", "", &http.Client{})
		assert.EqualError(t, err, "missing domain")
	})
}
