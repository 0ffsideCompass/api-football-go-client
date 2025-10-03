package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name          string
		query         string
		searchType    client.Type
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:       "successful search for teams",
			query:      "Manchester",
			searchType: client.Team,
			responseBody: `{
				"response": [
					{"team": {"id": 33, "name": "Manchester United", "country": "England"}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:       "successful search for players",
			query:      "Ronaldo",
			searchType: client.Player,
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
			name:       "successful search for leagues",
			query:      "Premier",
			searchType: client.League,
			responseBody: `{
				"response": [
					{"league": {"id": 39, "name": "Premier League", "country": "England"}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "search request with API error",
			query:         "InvalidQuery",
			searchType:    client.Team,
			responseBody:  `{"error": "Invalid search query"}`,
			statusCode:    http.StatusBadRequest,
			expectError:   true,
			errorContains: "API request failed with status 400",
		},
		{
			name:          "search request with 401 unauthorized",
			query:         "Manchester",
			searchType:    client.Team,
			responseBody:  `{"error": "Unauthorized"}`,
			statusCode:    http.StatusUnauthorized,
			expectError:   true,
			errorContains: "API request failed with status 401",
		},
		{
			name:          "search request with 403 forbidden",
			query:         "Manchester",
			searchType:    client.Team,
			responseBody:  `{"error": "Forbidden"}`,
			statusCode:    http.StatusForbidden,
			expectError:   true,
			errorContains: "API request failed with status 403",
		},
		{
			name:          "search request with 429 rate limit",
			query:         "Manchester",
			searchType:    client.Team,
			responseBody:  `{"error": "Rate limit exceeded"}`,
			statusCode:    http.StatusTooManyRequests,
			expectError:   true,
			errorContains: "API request failed with status 429",
		},
		{
			name:          "search request with 500 server error",
			query:         "Manchester",
			searchType:    client.Team,
			responseBody:  `{"error": "Internal server error"}`,
			statusCode:    http.StatusInternalServerError,
			expectError:   true,
			errorContains: "API request failed with status 500",
		},
		{
			name:       "search with empty results",
			query:      "NonExistentTeam",
			searchType: client.Team,
			responseBody: `{
				"response": [],
				"errors": [],
				"results": 0,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
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

			result, err := apiClient.Search(tt.query, tt.searchType)

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

func TestSearchTypes(t *testing.T) {
	// Test that the search type constants are properly defined
	tests := []struct {
		name       string
		searchType client.Type
	}{
		{
			name:       "team type",
			searchType: client.Team,
		},
		{
			name:       "player type",
			searchType: client.Player,
		},
		{
			name:       "league type",
			searchType: client.League,
		},
	}

	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"response": [], "errors": [], "results": 0}`)),
		},
	}

	apiClient, err := client.New("test-api-key", mockClient)
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that we can call search with each type without compilation errors
			_, err := apiClient.Search("test", tt.searchType)
			// We expect no error from the search type itself, any error should be from JSON parsing
			assert.NoError(t, err)
		})
	}
}
