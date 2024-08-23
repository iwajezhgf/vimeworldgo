package vimeworldgo

import "fmt"

// GuildSearch searches guilds by name or tag.
func (c *Client) GuildSearch(query string) ([]*Guild, error) {
	var guilds []*Guild

	q := fmt.Sprintf("?query=%s", query)
	if _, err := c.doRequest("GET", EndpointGuildSearch(q), &guilds); err != nil {
		return nil, err
	}

	return guilds, nil
}

// GetGuildById returns guild information by its id.
func (c *Client) GetGuildById(id int, unsafe ...bool) (*Guild, error) {
	query := fmt.Sprintf("?id=%d", id)
	return c.getGuild(query, unsafe...)
}

// GetGuildByName returns information about a guild by its name.
func (c *Client) GetGuildByName(name string, unsafe ...bool) (*Guild, error) {
	query := fmt.Sprintf("?name=%s", name)
	return c.getGuild(query, unsafe...)
}

// GetGuildByTag returns information about a guild by its tag.
func (c *Client) GetGuildByTag(tag string, unsafe ...bool) (*Guild, error) {
	query := fmt.Sprintf("?tag=%s", tag)
	return c.getGuild(query, unsafe...)
}

// Unsafe is true by default.
// https://vimeworld.github.io/api-docs/#apiguild_get_get
func (c *Client) getGuild(query string, unsafe ...bool) (*Guild, error) {
	var guild Guild

	isUnsafe := true
	if len(unsafe) > 0 {
		isUnsafe = unsafe[0]
	}

	if !isUnsafe {
		query += fmt.Sprintf("&unsafe=%t", unsafe)
	}

	if _, err := c.doRequest("GET", EndpointGuildGet(query), &guild); err != nil {
		return nil, err
	}

	return &guild, nil
}
