package vimeworldgo

// GetLeaderboardList returns list of all available record tables
func (c *Client) GetLeaderboardList() ([]*Leaderboard, error) {
	var leaderboards []*Leaderboard

	if _, err := c.doRequest("GET", EndpointLeaderboardList, &leaderboards); err != nil {
		return nil, err
	}

	return leaderboards, nil
}
