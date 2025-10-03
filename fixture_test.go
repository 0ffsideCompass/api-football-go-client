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

func TestFixturesLineups(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful fixtures lineups request",
			params: map[string]interface{}{"fixture": 12345},
			responseBody: `{
				"response": [
					{"team": {"id": 33, "name": "Manchester United"}, "formation": "4-3-3"}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "fixtures lineups request with API error",
			params:        map[string]interface{}{"fixture": -1},
			responseBody:  `{"error": "Invalid fixture ID"}`,
			statusCode:    http.StatusBadRequest,
			expectError:   true,
			errorContains: "API request failed with status 400",
		},
		{
			name:          "fixtures lineups request with invalid JSON response",
			params:        map[string]interface{}{"fixture": 12345},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling fixtures lineups response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.FixturesLineups(tt.params)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestFixturesEvents(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful fixtures events request",
			params: map[string]interface{}{"fixture": 12345},
			responseBody: `{
				"response": [
					{"time": {"elapsed": 15}, "team": {"id": 33}, "type": "Goal"}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "fixtures events request with invalid JSON response",
			params:        map[string]interface{}{"fixture": 12345},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling fixtures events response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.FixturesEvents(tt.params)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestFixture(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful fixture request",
			params: map[string]interface{}{"id": 12345},
			responseBody: `{
				"response": [
					{"fixture": {"id": 12345, "date": "2023-01-01T15:00:00Z"}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:   "fixture request with date range",
			params: map[string]interface{}{"from": "2023-01-01", "to": "2023-01-31"},
			responseBody: `{
				"response": [
					{"fixture": {"id": 12345, "date": "2023-01-15T15:00:00Z"}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "fixture request with invalid JSON response",
			params:        map[string]interface{}{"id": 12345},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling fixtures response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.Fixture(tt.params)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestFixtureHeadToHead(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful head to head request",
			params: map[string]interface{}{"h2h": "33-34"},
			responseBody: `{
				"response": [
					{"fixture": {"id": 12345}, "teams": {"home": {"id": 33}, "away": {"id": 34}}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "head to head request with invalid JSON response",
			params:        map[string]interface{}{"h2h": "33-34"},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling fixtures head to head response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.FixtureHeadToHead(tt.params)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestFixtureByDateAndLeague(t *testing.T) {
	fromDate := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2023, 1, 16, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name          string
		leagueID      int
		season        int
		fromDate      time.Time
		toDate        time.Time
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:     "successful fixture by date and league request",
			leagueID: 39,
			season:   2023,
			fromDate: fromDate,
			toDate:   toDate,
			responseBody: `{
				"response": [
					{"fixture": {"id": 12345, "date": "2023-01-15T15:00:00Z"}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "fixture by date and league request with invalid JSON response",
			leagueID:      39,
			season:        2023,
			fromDate:      fromDate,
			toDate:        toDate,
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling fixtures response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.FixtureByDateAndLeague(tt.leagueID, tt.season, tt.fromDate, tt.toDate)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestFixtureStatistics(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful fixture statistics request",
			params: map[string]interface{}{"fixture": 12345},
			responseBody: `{
				"response": [
					{"team": {"id": 33}, "statistics": [{"type": "Shots on Goal", "value": 5}]}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "fixture statistics request with invalid JSON response",
			params:        map[string]interface{}{"fixture": 12345},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling fixture statistics response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.FixtureStatistics(tt.params)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestFixturesPlayer(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful fixtures player request",
			params: map[string]interface{}{"fixture": 12345, "team": 33},
			responseBody: `{
				"response": [
					{"team": {"id": 33}, "players": [{"player": {"id": 276}}]}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "fixtures player request with invalid JSON response",
			params:        map[string]interface{}{"fixture": 12345, "team": 33},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling fixtures player response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.FixturesPlayer(tt.params)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}
