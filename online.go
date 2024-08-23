package vimeworldgo

// GetOnline returns the number of players online.
// In total and for each game separately.
func (c *Client) GetOnline() (*Online, error) {
	var online *Online

	if _, err := c.doRequest("GET", EndpointOnline, &online); err != nil {
		return nil, err
	}

	return online, nil
}

// GetOnlineStreams returns the list of streams that are currently running on the server.
func (c *Client) GetOnlineStreams() ([]*OnlineStreams, error) {
	var onlineStreams []*OnlineStreams

	if _, err := c.doRequest("GET", EndpointOnlineStreams, &onlineStreams); err != nil {
		return nil, err
	}

	return onlineStreams, nil
}

// GetOnlineStaff returns the list of online moderators.
func (c *Client) GetOnlineStaff() ([]*OnlineStaff, error) {
	var onlineStaff []*OnlineStaff

	if _, err := c.doRequest("GET", EndpointOnlineStaff, &onlineStaff); err != nil {
		return nil, err
	}

	return onlineStaff, nil
}
