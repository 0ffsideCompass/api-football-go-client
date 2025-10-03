// Package client provides a Go client for the Football API from API-Sports.io.
package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	domain = "https://v3.football.api-sports.io/"
)

const (
	authKey    = "X-RapidAPI-Key"
	hostKey    = "X-RapidAPI-Host"
	hostVal    = "api-football-v1.p.rapidapi.com"
	dateFormat = "2006-01-02"
)

// HttpClient is an interface that abstracts the http.Client's Do method,
// allowing for easier testing and customization.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client represents a client for the API-Football service.
// It holds the API key, the base domain URL, and the underlying HTTP client used to execute requests.
type Client struct {
	key    string
	Domain string
	client HttpClient
}

// New creates a new Client instance for the API-Football service using the default domain.
// It returns an error if the API key is missing or the provided HTTP client is nil.
func New(key string, client HttpClient) (*Client, error) {
	if key == "" {
		return nil, errors.New("missing key")
	}
	if client == nil {
		return nil, errors.New("missing http client")
	}
	return &Client{
		key:    key,
		Domain: domain,
		client: client,
	}, nil
}

// NewWithDomain creates a new Client instance with a custom domain.
// It returns an error if the API key is missing or the provided HTTP client is nil.
func NewWithDomain(key, domain string, client HttpClient) (*Client, error) {
	if key == "" {
		return nil, errors.New("missing key")
	}
	if client == nil {
		return nil, errors.New("missing http client")
	}
	return &Client{
		key:    key,
		Domain: domain,
		client: client,
	}, nil
}

func (c *Client) get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(authKey, c.key)
	req.Header.Add(hostKey, hostVal)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Check for non-2xx status codes
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", res.StatusCode, string(body))
	}

	return io.ReadAll(res.Body)
}

func (c *Client) formatDate(t time.Time) string {
	return t.Format(dateFormat)
}

// buildURL constructs the URL with dynamic parameters using net/url
func (c *Client) buildURL(endpoint string, params map[string]interface{}) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return endpoint
	}

	values := url.Values{}
	for key, value := range params {
		values.Add(key, fmt.Sprintf("%v", value))
	}

	u.RawQuery = values.Encode()
	return u.String()
}
