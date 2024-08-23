package vimeworldgo

// GetMiscGames returns a list of all games.
func (c *Client) GetMiscGames() ([]*Game, error) {
	var games []*Game

	if _, err := c.doRequest("GET", EndpointMiscGames, &games); err != nil {
		return nil, err
	}

	return games, nil
}

// GetMiscMaps returns a list of all maps on VimeWorld.
func (c *Client) GetMiscMaps() (*Maps, error) {
	var maps *Maps

	if _, err := c.doRequest("GET", EndpointMiscMaps, &maps); err != nil {
		return nil, err
	}

	return maps, nil
}

// GetMiscAchievements returns a list of all possible achievements with
// their descriptions, except secret achievements.
func (c *Client) GetMiscAchievements() (*Achievements, error) {
	var achievements *Achievements

	if _, err := c.doRequest("GET", EndpointMiscAchievements, &achievements); err != nil {
		return nil, err
	}

	return achievements, nil
}

// GetMiscToken returns information about the token.
func (c *Client) GetMiscToken(token string) (*Token, error) {
	var tkn *Token

	if _, err := c.doRequest("GET", EndpointMiscToken(token), &tkn); err != nil {
		return nil, err
	}

	return tkn, nil
}
