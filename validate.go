package client

import (
	"fmt"
	"strings"
)

// requireIntParams ensures that every given key is present in params and holds
// an integer value (an int, or a whole-number float64 as produced by JSON
// decoding).
func requireIntParams(params map[string]any, keys ...string) error {
	for _, key := range keys {
		value, ok := params[key]
		if !ok {
			return fmt.Errorf("'%s' is required", key)
		}
		if _, err := asInt(value); err != nil {
			return fmt.Errorf("'%s' must be an integer", key)
		}
	}
	return nil
}

// requireOneIntParam ensures that at least one of the given keys is present in
// params and that every provided key holds an integer value (an int, or a
// whole-number float64 as produced by JSON decoding).
func requireOneIntParam(params map[string]any, keys ...string) error {
	found := false
	for _, key := range keys {
		value, ok := params[key]
		if !ok {
			continue
		}
		found = true
		if _, err := asInt(value); err != nil {
			return fmt.Errorf("%s must be an integer", key)
		}
	}
	if !found {
		return fmt.Errorf("at least one of '%s' must be provided", strings.Join(keys, "' or '"))
	}
	return nil
}
