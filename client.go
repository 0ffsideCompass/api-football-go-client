package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	domain     = "https://v3.football.api-sports.io/"
	authKey    = "X-RapidAPI-Key"
	hostKey    = "X-RapidAPI-Host"
	hostVal    = "api-football-v1.p.rapidapi.com"
	dateFormat = "2006-01-02"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a struct that holds the key for the API
type Client struct {
	key    string
	Domain string
	client HttpClient
}

// NewClient creates a new client for the api-football service
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
	body, err := io.ReadAll(res.Body)

	fmt.Println(string(body))
	return body, err
}

func (c *Client) formatDate(t time.Time) string {
	return t.Format(dateFormat)
}

// buildURL constructs the URL with dynamic parameters using net/url
func (c *Client) buildURL(endpoint string, params map[string]interface{}) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err.Error()
	}

	values := url.Values{}
	for key, value := range params {
		values.Add(key, fmt.Sprintf("%v", value))
	}

	u.RawQuery = values.Encode()
	return u.String()
}
