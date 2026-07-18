package client

import (
	"fmt"
	"net/url"
)

// Type identifies the kind of entity a Search call looks up.
type Type string

// The entity kinds supported by Search.
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
// depending on the given Type, and returns the raw JSON response body.
// Unlike the other methods it does not unmarshal into a typed model.
//
// Note: the API requires search queries to be at least 3 characters
// (4 for players), and player searches may additionally require a league
// or team filter; such errors are reported in the response body.
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
