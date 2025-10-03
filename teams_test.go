package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func TestTeams(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful teams request",
			params: map[string]interface{}{"league": 39, "season": 2023},
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
			name:   "teams request with search parameter",
			params: map[string]interface{}{"search": "Manchester"},
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
			name:          "teams request with API error",
			params:        map[string]interface{}{"league": -1},
			responseBody:  `{"error": "Invalid league"}`,
			statusCode:    http.StatusBadRequest,
			expectError:   true,
			errorContains: "API request failed with status 400",
		},
		{
			name:          "teams request with invalid JSON response",
			params:        map[string]interface{}{"league": 39},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling teams response",
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

			result, err := apiClient.Teams(tt.params)

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

func TestTeamsStatistics(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful team statistics request",
			params: map[string]interface{}{"league": 39, "season": 2023, "team": 33},
			responseBody: `{
				"response": {
					"league": {"id": 39, "name": "Premier League"},
					"team": {"id": 33, "name": "Manchester United"}
				},
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "team statistics request with missing required params",
			params:        map[string]interface{}{"league": 39},
			responseBody:  `{"error": "Missing required parameters"}`,
			statusCode:    http.StatusBadRequest,
			expectError:   true,
			errorContains: "API request failed with status 400",
		},
		{
			name:          "team statistics request with invalid JSON response",
			params:        map[string]interface{}{"league": 39, "season": 2023, "team": 33},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling teams statistics response",
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

			result, err := apiClient.TeamsStatistics(tt.params)

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
