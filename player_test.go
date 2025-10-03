package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func TestPlayersSeasons(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful players seasons request",
			params: map[string]interface{}{"player": 276},
			responseBody: `{
				"response": [2020, 2021, 2022, 2023],
				"errors": [],
				"results": 4,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "players seasons request without player param",
			params:        map[string]interface{}{"season": 2023},
			responseBody:  ``,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "player must exist and be an integer",
		},
		{
			name:          "players seasons request with non-integer player param",
			params:        map[string]interface{}{"player": "invalid"},
			responseBody:  ``,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "player must exist and be an integer",
		},
		{
			name:          "players seasons request with API error",
			params:        map[string]interface{}{"player": -1},
			responseBody:  `{"error": "Invalid player ID"}`,
			statusCode:    http.StatusBadRequest,
			expectError:   true,
			errorContains: "API request failed with status 400",
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

			result, err := apiClient.PlayersSeasons(tt.params)

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

func TestPlayers(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful players request",
			params: map[string]interface{}{"team": 33, "season": 2023},
			responseBody: `{
				"response": [
					{"player": {"id": 276, "name": "Cristiano Ronaldo", "age": 38}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:   "players request with search parameter",
			params: map[string]interface{}{"search": "Ronaldo", "league": 39},
			responseBody: `{
				"response": [
					{"player": {"id": 276, "name": "Cristiano Ronaldo", "age": 38}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "players request with invalid JSON response",
			params:        map[string]interface{}{"team": 33},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling players response",
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

			result, err := apiClient.Players(tt.params)

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

func TestPlayersSquads(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful players squads request",
			params: map[string]interface{}{"team": 33, "player": 276},
			responseBody: `{
				"response": [
					{"team": {"id": 33, "name": "Manchester United"}, "players": []}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "players squads request without team param",
			params:        map[string]interface{}{"player": 276},
			responseBody:  ``,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "team must exist and be an integer",
		},
		{
			name:          "players squads request without player param",
			params:        map[string]interface{}{"team": 33},
			responseBody:  ``,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "player must exist and be an integer",
		},
		{
			name:          "players squads request with non-integer params",
			params:        map[string]interface{}{"team": "invalid", "player": "invalid"},
			responseBody:  ``,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "team must exist and be an integer",
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

			result, err := apiClient.PlayersSquads(tt.params)

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

func TestPlayersTopScorers(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful top scorers request",
			params: map[string]interface{}{"league": 39, "season": 2023},
			responseBody: `{
				"response": [
					{"player": {"id": 276, "name": "Cristiano Ronaldo"}, "statistics": [{"goals": {"total": 25}}]}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "top scorers request with invalid JSON response",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling players topscorers response",
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

			result, err := apiClient.PlayersTopScorers(tt.params)

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

func TestPlayersTopAssists(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful top assists request",
			params: map[string]interface{}{"league": 39, "season": 2023},
			responseBody: `{
				"response": [
					{"player": {"id": 276, "name": "Kevin De Bruyne"}, "statistics": [{"goals": {"assists": 15}}]}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "top assists request with invalid JSON response",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling players top assists response",
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

			result, err := apiClient.PlayersTopAssists(tt.params)

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

func TestPlayersTopCards(t *testing.T) {
	yellowCardTests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful top yellow cards request",
			params: map[string]interface{}{"league": 39, "season": 2023},
			responseBody: `{
				"response": [
					{"player": {"id": 276, "name": "Sergio Ramos"}, "statistics": [{"cards": {"yellow": 10}}]}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "top yellow cards request with invalid JSON response",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling players top yellow cards response",
		},
	}

	for _, tt := range yellowCardTests {
		t.Run("yellow_cards_"+tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.PlayersTopYellowCards(tt.params)

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

	redCardTests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful top red cards request",
			params: map[string]interface{}{"league": 39, "season": 2023},
			responseBody: `{
				"response": [
					{"player": {"id": 276, "name": "Sergio Ramos"}, "statistics": [{"cards": {"red": 2}}]}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "top red cards request with invalid JSON response",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling players top red cards response",
		},
	}

	for _, tt := range redCardTests {
		t.Run("red_cards_"+tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			result, err := apiClient.PlayersTopRedCards(tt.params)

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