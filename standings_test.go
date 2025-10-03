package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func TestStandings(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful standings request",
			params: map[string]interface{}{"league": 39, "season": 2023},
			responseBody: `{
				"response": [
					{
						"league": {"id": 39, "name": "Premier League"},
						"standings": [[
							{"rank": 1, "team": {"id": 50, "name": "Manchester City"}, "points": 89}
						]]
					}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:   "standings request with team filter",
			params: map[string]interface{}{"league": 39, "season": 2023, "team": 33},
			responseBody: `{
				"response": [
					{
						"league": {"id": 39, "name": "Premier League"},
						"standings": [[
							{"rank": 6, "team": {"id": 33, "name": "Manchester United"}, "points": 66}
						]]
					}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "standings request with API error",
			params:        map[string]interface{}{"league": -1, "season": 2023},
			responseBody:  `{"error": "Invalid league ID"}`,
			statusCode:    http.StatusBadRequest,
			expectError:   true,
			errorContains: "API request failed with status 400",
		},
		{
			name:          "standings request with invalid JSON response",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling standings response",
		},
		{
			name:   "standings request with nil params",
			params: nil,
			responseBody: `{
				"response": [],
				"errors": [],
				"results": 0,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "standings request with 401 unauthorized",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"error": "Unauthorized"}`,
			statusCode:    http.StatusUnauthorized,
			expectError:   true,
			errorContains: "API request failed with status 401",
		},
		{
			name:          "standings request with 403 forbidden",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"error": "Forbidden"}`,
			statusCode:    http.StatusForbidden,
			expectError:   true,
			errorContains: "API request failed with status 403",
		},
		{
			name:          "standings request with 404 not found",
			params:        map[string]interface{}{"league": 999999, "season": 2023},
			responseBody:  `{"error": "League not found"}`,
			statusCode:    http.StatusNotFound,
			expectError:   true,
			errorContains: "API request failed with status 404",
		},
		{
			name:          "standings request with 429 rate limit",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"error": "Rate limit exceeded"}`,
			statusCode:    http.StatusTooManyRequests,
			expectError:   true,
			errorContains: "API request failed with status 429",
		},
		{
			name:          "standings request with 500 server error",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"error": "Internal server error"}`,
			statusCode:    http.StatusInternalServerError,
			expectError:   true,
			errorContains: "API request failed with status 500",
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

			result, err := apiClient.Standings(tt.params)

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
