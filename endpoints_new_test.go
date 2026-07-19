package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
	"github.com/0ffsideCompass/api-football-go-client/models"
)

func newTestClient(t *testing.T, body string) *client.Client {
	t.Helper()
	mock := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		},
	}
	apiClient, err := client.New("test-api-key", mock)
	assert.NoError(t, err)
	return apiClient
}

// TestNewEndpointURLs verifies that the newer methods hit the endpoints they
// are named after.
func TestNewEndpointURLs(t *testing.T) {
	tests := []struct {
		name        string
		call        func(c *client.Client) error
		expectedURL string
	}{
		{
			name: "Timezone",
			call: func(c *client.Client) error {
				_, err := c.Timezone()
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/timezone",
		},
		{
			name: "Countries",
			call: func(c *client.Client) error {
				_, err := c.Countries(map[string]any{"name": "England"})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/countries?name=England",
		},
		{
			name: "LeaguesSeasons",
			call: func(c *client.Client) error {
				_, err := c.LeaguesSeasons()
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/leagues/seasons",
		},
		{
			name: "TeamsSeasons",
			call: func(c *client.Client) error {
				_, err := c.TeamsSeasons(map[string]any{"team": 33})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/teams/seasons?team=33",
		},
		{
			name: "TeamsCountries",
			call: func(c *client.Client) error {
				_, err := c.TeamsCountries()
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/teams/countries",
		},
		{
			name: "FixturesRounds",
			call: func(c *client.Client) error {
				_, err := c.FixturesRounds(map[string]any{"league": 39, "season": 2019})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/fixtures/rounds?league=39&season=2019",
		},
		{
			name: "Predictions",
			call: func(c *client.Client) error {
				_, err := c.Predictions(map[string]any{"fixture": 198772})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/predictions?fixture=198772",
		},
		{
			name: "PlayersProfiles",
			call: func(c *client.Client) error {
				_, err := c.PlayersProfiles(map[string]any{"player": 276})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/players/profiles?player=276",
		},
		{
			name: "PlayersTeams",
			call: func(c *client.Client) error {
				_, err := c.PlayersTeams(map[string]any{"player": 276})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/players/teams?player=276",
		},
		{
			name: "Odds",
			call: func(c *client.Client) error {
				_, err := c.Odds(map[string]any{"fixture": 326090, "bookmaker": 6})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/odds?bookmaker=6&fixture=326090",
		},
		{
			name: "OddsMapping",
			call: func(c *client.Client) error {
				_, err := c.OddsMapping(map[string]any{"page": 2})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/odds/mapping?page=2",
		},
		{
			name: "OddsBookmakers",
			call: func(c *client.Client) error {
				_, err := c.OddsBookmakers(nil)
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/odds/bookmakers",
		},
		{
			name: "OddsBets",
			call: func(c *client.Client) error {
				_, err := c.OddsBets(map[string]any{"search": "under"})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/odds/bets?search=under",
		},
		{
			name: "OddsLive",
			call: func(c *client.Client) error {
				_, err := c.OddsLive(map[string]any{"fixture": 721238})
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/odds/live?fixture=721238",
		},
		{
			name: "OddsLiveBets",
			call: func(c *client.Client) error {
				_, err := c.OddsLiveBets(nil)
				return err
			},
			expectedURL: "https://v3.football.api-sports.io/odds/live/bets",
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

func TestNewEndpointValidation(t *testing.T) {
	apiClient, err := client.New("test-api-key", newMockClient())
	assert.NoError(t, err)

	_, err = apiClient.TeamsSeasons(map[string]any{})
	assert.EqualError(t, err, "'team' is required")

	_, err = apiClient.FixturesRounds(map[string]any{"league": 39})
	assert.EqualError(t, err, "'season' is required")

	_, err = apiClient.Predictions(map[string]any{"fixture": "not-a-number"})
	assert.EqualError(t, err, "'fixture' must be an integer")

	_, err = apiClient.PlayersTeams(map[string]any{})
	assert.EqualError(t, err, "'player' is required")
}

func TestTimezoneUnmarshal(t *testing.T) {
	body := `{
		"get": "timezone", "parameters": [], "errors": [], "results": 2,
		"paging": {"current": 1, "total": 1},
		"response": ["Africa/Abidjan", "Europe/London"]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.Timezone()
	assert.NoError(t, err)
	assert.Equal(t, []string{"Africa/Abidjan", "Europe/London"}, resp.Response)
}

func TestCountriesUnmarshal(t *testing.T) {
	body := `{
		"get": "countries", "parameters": {"name": "england"}, "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{"name": "England", "code": "GB", "flag": "https://media.api-sports.io/flags/gb.svg"}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.Countries(map[string]any{"name": "england"})
	assert.NoError(t, err)
	assert.Len(t, resp.Response, 1)
	assert.Equal(t, "England", resp.Response[0].Name)
	assert.Equal(t, "GB", resp.Response[0].Code)
}

func TestSeasonsUnmarshal(t *testing.T) {
	body := `{
		"get": "leagues/seasons", "parameters": [], "errors": [], "results": 3,
		"paging": {"current": 1, "total": 1},
		"response": [2018, 2019, 2020]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.LeaguesSeasons()
	assert.NoError(t, err)
	assert.Equal(t, []int{2018, 2019, 2020}, resp.Response)
}

func TestFixturesRoundsUnmarshal(t *testing.T) {
	t.Run("plain string rounds", func(t *testing.T) {
		body := `{
			"get": "fixtures/rounds", "parameters": {}, "errors": [], "results": 2,
			"paging": {"current": 1, "total": 1},
			"response": ["Regular Season - 1", "Regular Season - 2"]
		}`
		apiClient := newTestClient(t, body)

		resp, err := apiClient.FixturesRounds(map[string]any{"league": 39, "season": 2019})
		assert.NoError(t, err)
		assert.Len(t, resp.Response, 2)
		assert.Equal(t, "Regular Season - 1", resp.Response[0].Round)
		assert.Empty(t, resp.Response[0].Dates)
	})

	t.Run("rounds with dates", func(t *testing.T) {
		body := `{
			"get": "fixtures/rounds", "parameters": {}, "errors": [], "results": 1,
			"paging": {"current": 1, "total": 1},
			"response": [{"round": "Regular Season - 1", "dates": ["2024-08-16", "2024-08-17"]}]
		}`
		apiClient := newTestClient(t, body)

		resp, err := apiClient.FixturesRounds(map[string]any{"league": 39, "season": 2024, "dates": "true"})
		assert.NoError(t, err)
		assert.Len(t, resp.Response, 1)
		assert.Equal(t, "Regular Season - 1", resp.Response[0].Round)
		assert.Equal(t, []string{"2024-08-16", "2024-08-17"}, resp.Response[0].Dates)
	})
}

func TestPredictionsUnmarshal(t *testing.T) {
	body := `{
		"get": "predictions", "parameters": {"fixture": "198772"}, "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{
			"predictions": {
				"winner": {"id": 1189, "name": "Deportivo Santani", "comment": "Win or draw"},
				"win_or_draw": true,
				"under_over": "-3.5",
				"goals": {"home": "-2.5", "away": "-1.5"},
				"advice": "Combo Double chance : Deportivo Santani or draw and -3.5 goals",
				"percent": {"home": "45%", "draw": "45%", "away": "10%"}
			},
			"league": {"id": 252, "name": "Primera Division - Clausura", "country": "Paraguay", "logo": "l", "flag": "f", "season": 2019},
			"teams": {
				"home": {
					"id": 1189, "name": "Deportivo Santani", "logo": "l",
					"last_5": {
						"form": "60%", "att": "60%", "def": "0%",
						"goals": {"for": {"total": 3, "average": 0.6}, "against": {"total": 5, "average": 1}}
					},
					"league": {
						"form": "LDLDLDLWLWWLW",
						"fixtures": {
							"played": {"home": 6, "away": 7, "total": 13},
							"wins": {"home": 2, "away": 2, "total": 4},
							"draws": {"home": 3, "away": 0, "total": 3},
							"loses": {"home": 1, "away": 5, "total": 6}
						},
						"goals": {
							"for": {"total": {"home": 7, "away": 4, "total": 11}, "average": {"home": "1.2", "away": "0.6", "total": "0.8"}},
							"against": {"total": {"home": 6, "away": 14, "total": 20}, "average": {"home": "1.0", "away": "2.0", "total": "1.5"}}
						},
						"biggest": {
							"streak": {"wins": 2, "draws": 1, "loses": 1},
							"wins": {"home": "3-1", "away": "0-1"},
							"loses": {"home": "0-2", "away": "4-0"},
							"goals": {"for": {"home": 3, "away": 1}, "against": {"home": 2, "away": 4}}
						},
						"clean_sheet": {"home": 1, "away": 2, "total": 3},
						"failed_to_score": {"home": 1, "away": 3, "total": 4}
					}
				},
				"away": {"id": 1180, "name": "Deportivo Capiata", "logo": "l"}
			},
			"comparison": {
				"form": {"home": "60%", "away": "40%"},
				"att": {"home": "43%", "away": "57%"},
				"def": {"home": "62%", "away": "38%"},
				"poisson_distribution": {"home": "75%", "away": "25%"},
				"h2h": {"home": "29%", "away": "71%"},
				"goals": {"home": "40%", "away": "60%"},
				"total": {"home": "51.5%", "away": "48.5%"}
			},
			"h2h": [{
				"fixture": {"id": 198706, "referee": "J. Mendez", "timezone": "UTC", "date": "2019-07-27T19:30:00+00:00", "timestamp": 1564255800},
				"goals": {"home": 3, "away": 1}
			}]
		}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.Predictions(map[string]any{"fixture": 198772})
	assert.NoError(t, err)
	assert.Len(t, resp.Response, 1)
	item := resp.Response[0]
	assert.Equal(t, "Deportivo Santani", item.Predictions.Winner.Name)
	assert.True(t, item.Predictions.WinOrDraw)
	assert.Equal(t, "45%", item.Predictions.Percent.Home)
	assert.Equal(t, 13, item.Teams.Home.League.Fixtures.Played.Total)
	assert.Equal(t, "1.2", item.Teams.Home.League.Goals.For.Average.Home)
	assert.Equal(t, models.FlexString("0.6"), item.Teams.Home.Last5.Goals.For.Average)
	assert.Equal(t, "75%", item.Comparison.PoissonDistribution.Home)
	assert.Len(t, item.H2H, 1)
	assert.Equal(t, 198706, item.H2H[0].Fixture.ID)
}

func TestPlayersProfilesUnmarshal(t *testing.T) {
	body := `{
		"get": "players/profiles", "parameters": {"player": "276"}, "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{
			"player": {
				"id": 276, "name": "Neymar", "firstname": "Neymar", "lastname": "da Silva Santos Junior",
				"age": 32,
				"birth": {"date": "1992-02-05", "place": "Mogi das Cruzes", "country": "Brazil"},
				"nationality": "Brazil", "height": "175 cm", "weight": "68 kg",
				"number": 10, "position": "Attacker", "photo": "p"
			}
		}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.PlayersProfiles(map[string]any{"player": 276})
	assert.NoError(t, err)
	assert.Len(t, resp.Response, 1)
	assert.Equal(t, "Neymar", resp.Response[0].Player.Name)
	assert.Equal(t, 10, resp.Response[0].Player.Number)
	assert.Equal(t, "Brazil", resp.Response[0].Player.Birth.Country)
}

func TestPlayersTeamsUnmarshal(t *testing.T) {
	body := `{
		"get": "players/teams", "parameters": {"player": "276"}, "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{
			"team": {"id": 85, "name": "Paris Saint Germain", "logo": "l"},
			"seasons": [2022, 2021, 2020]
		}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.PlayersTeams(map[string]any{"player": 276})
	assert.NoError(t, err)
	assert.Len(t, resp.Response, 1)
	assert.Equal(t, "Paris Saint Germain", resp.Response[0].Team.Name)
	assert.Equal(t, []int{2022, 2021, 2020}, resp.Response[0].Seasons)
}

func TestOddsUnmarshal(t *testing.T) {
	body := `{
		"get": "odds", "parameters": {"fixture": "326090"}, "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{
			"league": {"id": 116, "name": "Vysshaya Liga", "country": "Belarus", "logo": "l", "flag": "f", "season": 2020},
			"fixture": {"id": 326090, "timezone": "UTC", "date": "2020-05-15T15:00:00+00:00", "timestamp": 1589554800},
			"update": "2020-05-15T09:49:32+00:00",
			"bookmakers": [{
				"id": 6, "name": "Bwin",
				"bets": [{
					"id": 38, "name": "Exact Goals Number",
					"values": [
						{"value": 4, "odd": "7.00"},
						{"value": "more 8", "odd": "251.00"}
					]
				}]
			}]
		}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.Odds(map[string]any{"fixture": 326090})
	assert.NoError(t, err)
	assert.Len(t, resp.Response, 1)
	item := resp.Response[0]
	assert.Equal(t, 326090, item.Fixture.ID)
	assert.Equal(t, "Bwin", item.Bookmakers[0].Name)
	values := item.Bookmakers[0].Bets[0].Values
	// The API mixes numeric and string selection labels.
	assert.Equal(t, float64(4), values[0].Value)
	assert.Equal(t, "7.00", values[0].Odd)
	assert.Equal(t, "more 8", values[1].Value)
}

func TestOddsMappingUnmarshal(t *testing.T) {
	body := `{
		"get": "odds/mapping", "parameters": [], "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{
			"league": {"id": 106, "season": 2019},
			"fixture": {"id": 154507, "date": "2020-05-29T18:30:00+00:00", "timestamp": 1590777000},
			"update": "2020-05-15T09:52:28+00:00"
		}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.OddsMapping(nil)
	assert.NoError(t, err)
	assert.Len(t, resp.Response, 1)
	assert.Equal(t, 154507, resp.Response[0].Fixture.ID)
	assert.Equal(t, 2019, resp.Response[0].League.Season)
}

func TestOddsBookmakersAndBetsUnmarshal(t *testing.T) {
	body := `{
		"get": "odds/bookmakers", "parameters": [], "errors": [], "results": 2,
		"paging": {"current": 1, "total": 1},
		"response": [{"id": 6, "name": "Bwin"}, {"id": 8, "name": "Bet365"}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.OddsBookmakers(nil)
	assert.NoError(t, err)
	assert.Len(t, resp.Response, 2)
	assert.Equal(t, "Bet365", resp.Response[1].Name)

	betsBody := `{
		"get": "odds/bets", "parameters": {"search": "under"}, "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{"id": 5, "name": "Goals Over/Under"}]
	}`
	apiClient = newTestClient(t, betsBody)

	bets, err := apiClient.OddsBets(map[string]any{"search": "under"})
	assert.NoError(t, err)
	assert.Equal(t, "Goals Over/Under", bets.Response[0].Name)
}

func TestOddsLiveUnmarshal(t *testing.T) {
	body := `{
		"get": "odds/live", "parameters": {"fixture": "721238"}, "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{
			"fixture": {"id": 721238, "status": {"long": "Second Half", "elapsed": 62, "seconds": "62:14"}},
			"league": {"id": 30, "season": 2022},
			"teams": {"home": {"id": 1563, "goals": 1}, "away": {"id": 1565, "goals": 0}},
			"status": {"stopped": false, "blocked": false, "finished": false},
			"update": "2022-01-27T16:21:01+00:00",
			"odds": [{
				"id": 20, "name": "Match Corners",
				"values": [
					{"value": "Over", "odd": "2.5", "handicap": "8", "main": null, "suspended": false}
				]
			}]
		}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.OddsLive(map[string]any{"fixture": 721238})
	assert.NoError(t, err)
	assert.Len(t, resp.Response, 1)
	item := resp.Response[0]
	assert.Equal(t, 721238, item.Fixture.ID)
	assert.Equal(t, "62:14", item.Fixture.Status.Seconds)
	assert.Equal(t, 1, item.Teams.Home.Goals)
	assert.False(t, item.Status.Stopped)
	assert.Equal(t, "Match Corners", item.Odds[0].Name)
	assert.Equal(t, "Over", item.Odds[0].Values[0].Value)
	assert.Equal(t, "8", item.Odds[0].Values[0].Handicap)
	assert.False(t, item.Odds[0].Values[0].Suspended)
}
