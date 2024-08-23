package vimeworldgo

import (
	"fmt"
	"strings"
)

// GetLocale returns localized names of games, statistics, and ranks.
func (c *Client) GetLocale(lang string, parts ...string) (*Locale, error) {
	var locale *Locale
	var query string

	if len(parts) > 0 {
		query = fmt.Sprintf("?parts=%s", strings.Join(parts, ","))
	}

	if _, err := c.doRequest("GET", EndpointLocaleName(lang, query), &locale); err != nil {
		return nil, err
	}

	return locale, nil
}
