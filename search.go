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

func (c *Client) Search(q string, t Type) ([]byte, error) {
	// Correctly construct the path
	path := fmt.Sprintf("%ss/%s/", t.String(), t.String())

	// Now escape the query and append it to the path
	escapedQuery := url.PathEscape(q)
	fullURL := c.Domain + path + escapedQuery

	fmt.Println(fullURL)
	body, err := c.get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error getting search results: %w", err)
	}

	return body, nil
}
