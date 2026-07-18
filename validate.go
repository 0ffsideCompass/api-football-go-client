package client

import (
	"fmt"
	"strings"
)

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
