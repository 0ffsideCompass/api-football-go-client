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

func TestBuildURL(t *testing.T) {
	// buildURL is private, so exercise it through Leagues and inspect the
	// query string of the request the client actually sent.
	tests := []struct {
		name          string
		params        map[string]any
		expectedQuery string
	}{
		{
			name:          "no params",
			params:        nil,
			expectedQuery: "",
		},
		{
			name:          "single param",
			params:        map[string]any{"league": 39},
			expectedQuery: "league=39",
		},
		{
			name: "multiple params are sorted by key",
			params: map[string]any{
				"team":   33,
				"league": 39,
				"season": 2023,
			},
			expectedQuery: "league=39&season=2023&team=33",
		},
		{
			name: "string params are escaped",
			params: map[string]any{
				"search": "Manchester United",
			},
			expectedQuery: "search=Manchester+United",
		},
		{
			name: "whole float64 params render as integers, not scientific notation",
			params: map[string]any{
				"id": float64(1581037),
			},
			expectedQuery: "id=1581037",
		},
		{
			name: "fractional float64 params keep their decimals",
			params: map[string]any{
				"x": 2.5,
			},
			expectedQuery: "x=2.5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := newMockClient()
			apiClient, err := client.New("test-key", mockClient)
			assert.NoError(t, err)

			_, err = apiClient.Leagues(tt.params)
			assert.NoError(t, err)
			assert.NotNil(t, mockClient.LastRequest)
			assert.Equal(t, tt.expectedQuery, mockClient.LastRequest.URL.RawQuery)
		})
	}
}

func TestFormatDate(t *testing.T) {
	// Test the date formatting functionality indirectly through methods that use it
	testDate := time.Date(2023, 1, 15, 10, 30, 0, 0, time.UTC)

	// Create a mock client with a proper response structure for fixtures
	httpClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"response": [], "errors": {}, "results": 0, "paging": {"current": 1, "total": 1}}`)),
		},
	}
	apiClient, err := client.New("test-key", httpClient)
	assert.NoError(t, err)

	// Test that date formatting works correctly through FixtureByDateAndLeague
	// which uses the formatDate method internally
	testDateTo := time.Date(2023, 1, 16, 10, 30, 0, 0, time.UTC)
	_, err = apiClient.FixtureByDateAndLeague(39, 2023, testDate, testDateTo)

	// The error should be from HTTP/JSON parsing, not from date formatting
	// If date formatting was broken, we'd get a different type of error
	assert.NoError(t, err) // Should succeed with valid JSON response
}

func TestClientConfiguration(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		domain        string
		expectError   bool
		errorContains string
	}{
		{
			name:        "successful client creation with default domain",
			key:         "valid-api-key",
			expectError: false,
		},
		{
			name:        "successful client creation with custom domain",
			key:         "valid-api-key",
			domain:      "https://custom.api.com/",
			expectError: false,
		},
		{
			name:          "client creation with empty key",
			key:           "",
			expectError:   true,
			errorContains: "missing key",
		},
	}

	httpClient := &MockHTTPClient{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var apiClient *client.Client
			var err error

			if tt.domain != "" {
				apiClient, err = client.NewWithDomain(tt.key, tt.domain, httpClient)
			} else {
				apiClient, err = client.New(tt.key, httpClient)
			}

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, apiClient)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, apiClient)
			}
		})
	}
}

func TestClientWithNilHTTPClient(t *testing.T) {
	// Test that client creation fails with nil HTTP client
	apiClient, err := client.New("test-key", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing http client")
	assert.Nil(t, apiClient)

	// Test NewWithDomain with nil HTTP client
	apiClient, err = client.NewWithDomain("test-key", "https://custom.com", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing http client")
	assert.Nil(t, apiClient)
}
