package client

import (
	"fmt"
	"net/url"
)

type Type string

const (
	Team   Type = "team"
	League Type = "league"
	Player Type = "player"
)

// String method to return the string representation of Type
func (t Type) String() string {
	return string(t)
}

// Search queries the /teams, /leagues, or /players endpoint by name,
// depending on the given Type, and returns the raw response body.
func (c *Client) Search(q string, t Type) ([]byte, error) {
	values := url.Values{}
	values.Set("search", q)
	fullURL := fmt.Sprintf("%s%ss?%s", c.Domain, t.String(), values.Encode())

	body, err := c.get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error getting search results: %w", err)
	}

	return body, nil
}
