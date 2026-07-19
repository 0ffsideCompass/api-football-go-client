package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/0ffsideCompass/api-football-go-client/models"
)

// Regression test for a production failure: API-Football mixes value types in
// statistics arrays ("55%", 12, null), which previously failed to unmarshal
// into string- or int-typed Value fields.

func TestFixtureStatisticsMixedValues(t *testing.T) {
	body := `{
		"get": "fixtures/statistics", "parameters": {"fixture": "1581037"}, "errors": [], "results": 2,
		"paging": {"current": 1, "total": 1},
		"response": [{
			"team": {"id": 33, "name": "Manchester United", "logo": "l"},
			"statistics": [
				{"type": "Ball Possession", "value": "55%"},
				{"type": "Total Shots", "value": 12},
				{"type": "expected_goals", "value": "1.29"},
				{"type": "Yellow Cards", "value": null}
			]
		}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.FixtureStatistics(map[string]any{"fixture": 1581037})
	assert.NoError(t, err)
	stats := resp.Response[0].Statistics
	assert.Equal(t, models.FlexString("55%"), stats[0].Value)
	assert.Equal(t, models.FlexString("12"), stats[1].Value)
	assert.Equal(t, models.FlexString("1.29"), stats[2].Value)
	assert.Equal(t, models.FlexString(""), stats[3].Value)
}

func TestFixtureWithMixedStatisticsValues(t *testing.T) {
	// The /fixtures endpoint embeds the same statistics shape when queried by id.
	body := `{
		"get": "fixtures", "parameters": {"id": "1581037"}, "errors": [], "results": 1,
		"paging": {"current": 1, "total": 1},
		"response": [{
			"fixture": {"id": 1581037},
			"statistics": [{
				"team": {"id": 33, "name": "Manchester United", "logo": "l"},
				"statistics": [
					{"type": "Ball Possession", "value": "55%"},
					{"type": "Total Shots", "value": 12},
					{"type": "Yellow Cards", "value": null}
				]
			}]
		}]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.Fixture(map[string]any{"id": 1581037})
	assert.NoError(t, err)
	stats := resp.Response[0].Statistics[0].Statistics
	assert.Equal(t, models.FlexString("55%"), stats[0].Value)
	assert.Equal(t, models.FlexString("12"), stats[1].Value)
	assert.Equal(t, models.FlexString(""), stats[2].Value)
}

func TestFixturesEventsStoppageTime(t *testing.T) {
	// The API reports stoppage-time events with a numeric "extra" (e.g. 90+5)
	// and null otherwise; both must decode.
	body := `{
		"get": "fixtures/events", "parameters": {"fixture": "1581037"}, "errors": [], "results": 2,
		"paging": {"current": 1, "total": 1},
		"response": [
			{"time": {"elapsed": 25, "extra": null}, "type": "Goal", "detail": "Normal Goal"},
			{"time": {"elapsed": 90, "extra": 5}, "type": "Card", "detail": "Yellow Card"}
		]
	}`
	apiClient := newTestClient(t, body)

	resp, err := apiClient.FixturesEvents(map[string]any{"fixture": 1581037})
	assert.NoError(t, err)
	assert.Equal(t, models.FlexString(""), resp.Response[0].Time.Extra)
	assert.Equal(t, 90, resp.Response[1].Time.Elapsed)
	assert.Equal(t, models.FlexString("5"), resp.Response[1].Time.Extra)
}

func TestFlexStringRejectsInvalidJSON(t *testing.T) {
	var f models.FlexString
	assert.Error(t, f.UnmarshalJSON([]byte(`{"not": "a scalar"}`)))
}
