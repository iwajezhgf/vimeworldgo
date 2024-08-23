package vimeworldgo

import (
	"fmt"
	"strconv"
	"strings"
)

// GetUsersByNames returns information about a player or multiple players by their username.
func (c *Client) GetUsersByNames(names ...string) ([]*User, error) {
	var users []*User

	n := strings.Join(names, ",")
	if _, err := c.doRequest("GET", EndpointUserNames(n), &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUsersByIds returns information about a player or multiple players by their id.
func (c *Client) GetUsersByIds(ids ...int) ([]*User, error) {
	var users []*User

	strIds := strings.Join(i2s(ids), ",")
	if _, err := c.doRequest("GET", EndpointUserIds(strIds), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func i2s(ids []int) []string {
	strIds := make([]string, len(ids))
	for i, id := range ids {
		strIds[i] = strconv.Itoa(id)
	}
	return strIds
}

// GetUserFriends returns a list of the player's friends.
func (c *Client) GetUserFriends(id int) (*Friends, error) {
	var friends *Friends

	if _, err := c.doRequest("GET", EndpointUserFriends(strconv.Itoa(id)), &friends); err != nil {
		return nil, err
	}

	return friends, nil
}

// GetUserSession returns the online player status.
func (c *Client) GetUserSession(id int) (*Session, error) {
	var session *Session

	if _, err := c.doRequest("GET", EndpointUserSession(strconv.Itoa(id)), &session); err != nil {
		return nil, err
	}

	return session, nil
}

// GetUserStats returns the statistics of all games of the requested player.
func (c *Client) GetUserStats(id int, games ...string) (*UserStats, error) {
	var stats *UserStats
	var query string

	if len(games) > 0 {
		query = fmt.Sprintf("?games=%s", strings.Join(games, ","))
	}

	if _, err := c.doRequest("GET", EndpointUserStats(strconv.Itoa(id), query), &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// GetUserAchievements returns a list of all player's achievements.
func (c *Client) GetUserAchievements(id int) (*UserAchievements, error) {
	var achievements *UserAchievements

	if _, err := c.doRequest("GET", EndpointUserAchievements(strconv.Itoa(id)), &achievements); err != nil {
		return nil, err
	}

	return achievements, nil
}

// GetUserLeaderboards returns a player's place in the tops.
func (c *Client) GetUserLeaderboards(id int) (*UserLeaderboards, error) {
	var leaderboards UserLeaderboards

	if _, err := c.doRequest("GET", EndpointUserLeaderboards(strconv.Itoa(id)), &leaderboards); err != nil {
		return nil, err
	}

	return &leaderboards, nil
}

// GetUserMatchesAfter returns the user matches after the specified match id.
func (c *Client) GetUserMatchesAfter(id int, after string, count int) (*UserMatches, error) {
	query := fmt.Sprintf("?after=%s", after)
	return c.getUserMatches(id, query, count)
}

// GetUserMatchesBefore returns the user matches before the specified match id.
func (c *Client) GetUserMatchesBefore(id int, before string, count int) (*UserMatches, error) {
	query := fmt.Sprintf("?before=%s", before)
	return c.getUserMatches(id, query, count)
}

// GetUserMatchesOffset returns the user matches with offset.
func (c *Client) GetUserMatchesOffset(id, offset, count int) (*UserMatches, error) {
	query := fmt.Sprintf("?offset=%d", offset)
	return c.getUserMatches(id, query, count)
}

func (c *Client) getUserMatches(id int, query string, count int) (*UserMatches, error) {
	var matches *UserMatches

	if count > 0 {
		query += fmt.Sprintf("&count=%d", count)
	}

	if _, err := c.doRequest("GET", EndpointUserMatches(strconv.Itoa(id), query), &matches); err != nil {
		return nil, err
	}

	return matches, nil
}
