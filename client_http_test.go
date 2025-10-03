package client_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

// MockHTTPClient is a mock implementation of the HttpClient interface
type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Response, nil
}

func TestHTTPStatusValidation(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   string
		expectError    bool
		errorContains  string
	}{
		{
			name:         "should succeed with 200 OK",
			statusCode:   http.StatusOK,
			responseBody: `{"response":[],"errors":[],"results":0,"paging":{"current":1,"total":1}}`,
			expectError:  false,
		},
		{
			name:          "should fail with 400 Bad Request",
			statusCode:    http.StatusBadRequest,
			responseBody:  `{"error":"Bad Request"}`,
			expectError:   true,
			errorContains: "API request failed with status 400",
		},
		{
			name:          "should fail with 401 Unauthorized",
			statusCode:    http.StatusUnauthorized,
			responseBody:  `{"error":"Unauthorized"}`,
			expectError:   true,
			errorContains: "API request failed with status 401",
		},
		{
			name:          "should fail with 403 Forbidden",
			statusCode:    http.StatusForbidden,
			responseBody:  `{"error":"Forbidden"}`,
			expectError:   true,
			errorContains: "API request failed with status 403",
		},
		{
			name:          "should fail with 404 Not Found",
			statusCode:    http.StatusNotFound,
			responseBody:  `{"error":"Not Found"}`,
			expectError:   true,
			errorContains: "API request failed with status 404",
		},
		{
			name:          "should fail with 429 Too Many Requests",
			statusCode:    http.StatusTooManyRequests,
			responseBody:  `{"error":"Rate limit exceeded"}`,
			expectError:   true,
			errorContains: "API request failed with status 429",
		},
		{
			name:          "should fail with 500 Internal Server Error",
			statusCode:    http.StatusInternalServerError,
			responseBody:  `{"error":"Internal Server Error"}`,
			expectError:   true,
			errorContains: "API request failed with status 500",
		},
		{
			name:          "should fail with 503 Service Unavailable",
			statusCode:    http.StatusServiceUnavailable,
			responseBody:  `{"error":"Service Unavailable"}`,
			expectError:   true,
			errorContains: "API request failed with status 503",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock HTTP client
			mockClient := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: tt.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.responseBody)),
				},
			}

			// Create API client
			apiClient, err := client.New("test-api-key", mockClient)
			assert.NoError(t, err)

			// Test with Leagues endpoint as an example
			_, err = apiClient.Leagues(nil)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			} else {
				// The error might be from JSON unmarshalling for empty responses,
				// but not from HTTP status check
				if err != nil {
					assert.NotContains(t, err.Error(), "API request failed with status")
				}
			}
		})
	}
}