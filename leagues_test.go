package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func TestLeagues(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]interface{}
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
	}{
		{
			name:   "successful leagues request",
			params: map[string]interface{}{"country": "England"},
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
			name:   "leagues request with search parameter",
			params: map[string]interface{}{"search": "Premier League"},
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
			name:          "leagues request with API error",
			params:        map[string]interface{}{"country": "InvalidCountry"},
			responseBody:  `{"error": "Invalid country"}`,
			statusCode:    http.StatusBadRequest,
			expectError:   true,
			errorContains: "API request failed with status 400",
		},
		{
			name:   "leagues request with nil params",
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
			name:          "leagues request with invalid JSON response",
			params:        map[string]interface{}{"country": "England"},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling leagues response",
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

			result, err := apiClient.Leagues(tt.params)

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
