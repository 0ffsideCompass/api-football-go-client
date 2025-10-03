package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func TestInjuries(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful injuries request",
			params: map[string]interface{}{"league": 39, "season": 2023},
			responseBody: `{
				"response": [
					{"player": {"id": 276, "name": "Player Name"}, "team": {"id": 33}, "fixture": {"id": 12345}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "injuries request with invalid JSON response",
			params:        map[string]interface{}{"league": 39, "season": 2023},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling injuries response",
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

			result, err := apiClient.Injuries(tt.params)

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

func TestCoachs(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful coaches request",
			params: map[string]interface{}{"team": 33},
			responseBody: `{
				"response": [
					{"id": 1, "name": "Coach Name", "team": {"id": 33, "name": "Manchester United"}}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "coaches request with invalid JSON response",
			params:        map[string]interface{}{"team": 33},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling coachs response",
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

			result, err := apiClient.Coachs(tt.params)

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

func TestSidelined(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful sidelined request",
			params: map[string]interface{}{"player": 276},
			responseBody: `{
				"response": [
					{"type": "Injury", "start": "2023-01-01", "end": "2023-01-15"}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "sidelined request with invalid JSON response",
			params:        map[string]interface{}{"player": 276},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling sidelined response",
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

			result, err := apiClient.Sidelined(tt.params)

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

func TestTransfers(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful transfers request",
			params: map[string]interface{}{"player": 276},
			responseBody: `{
				"response": [
					{"player": {"id": 276}, "transfers": [{"date": "2023-01-01", "teams": {"in": {"id": 33}, "out": {"id": 34}}}]}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "transfers request with invalid JSON response",
			params:        map[string]interface{}{"player": 276},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling transfers response",
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

			result, err := apiClient.Transfers(tt.params)

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

func TestTrophies(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful trophies request",
			params: map[string]interface{}{"player": 276},
			responseBody: `{
				"response": [
					{"league": "Premier League", "country": "England", "season": "2022-2023", "place": "Winner"}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "trophies request with invalid JSON response",
			params:        map[string]interface{}{"player": 276},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling trophies response",
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

			result, err := apiClient.Trophies(tt.params)

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

func TestVenues(t *testing.T) {
	tests := []struct {
		name           string
		params         map[string]interface{}
		responseBody   string
		statusCode     int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful venues request",
			params: map[string]interface{}{"city": "Manchester"},
			responseBody: `{
				"response": [
					{"id": 1, "name": "Old Trafford", "city": "Manchester", "capacity": 76000}
				],
				"errors": [],
				"results": 1,
				"paging": {"current": 1, "total": 1}
			}`,
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:          "venues request with invalid JSON response",
			params:        map[string]interface{}{"city": "Manchester"},
			responseBody:  `{"invalid": json}`,
			statusCode:    http.StatusOK,
			expectError:   true,
			errorContains: "error unmarshalling venues response",
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

			result, err := apiClient.Venues(tt.params)

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