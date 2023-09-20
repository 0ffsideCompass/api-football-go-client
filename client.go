package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	domain  = "https://v3.football.api-sports.io/"
	authKey = "X-RapidAPI-Key"
	hostKey = "X-RapidAPI-Host"
	hostVal = "api-football-v1.p.rapidapi.com"
)

// Client is a struct that holds the key for the API
type Client struct {
	key    string
	domain string
	client *http.Client
}

// NewClient creates a new client for the api-football service
func New(key string) (*Client, error) {
	if key == "" {
		return nil, errors.New("missing key")
	}
	return &Client{
		key:    key,
		domain: domain,
		client: &http.Client{},
	}, nil
}

func (c *Client) get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.domain, endpoint)
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
	body, err := io.ReadAll(res.Body)

	return body, err
}
