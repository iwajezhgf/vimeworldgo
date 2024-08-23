package vimeworldgo

// User represents a user.
type User struct {
	ID              int      `json:"id"`
	Username        string   `json:"username"`
	Level           int      `json:"level"`
	LevelPercentage float64  `json:"levelPercentage"`
	Rank            string   `json:"rank"`
	Ranks           []string `json:"ranks"`
	PlayedSeconds   int64    `json:"playedSeconds"`
	LastSeen        int64    `json:"lastSeen"`
	Prime           bool     `json:"prime"`
	PrimeStart      int64    `json:"primeStart"`
	PrimeIcon       string   `json:"primeIcon"`
	Prefix          string   `json:"prefix"`
	CustomColors    []string `json:"customColors"`

	// Guild represents the user's current guild membership.
	// It may be nil if the user is not part of any guild.
	Guild *GuildBrief `json:"guild,omitempty"`
}

// Friends represents a user's friends list.
type Friends struct {
	User    User   `json:"user"`
	Friends []User `json:"friends"`
}

// Session represents a user's session data.
type Session struct {
	User   User          `json:"user"`
	Online OnlineSession `json:"online"`
}

// OnlineSession represents the online status of a user.
type OnlineSession struct {
	Value   bool   `json:"value"`
	Message string `json:"message"`

	// Game is the name of the game the user is currently playing, if any.
	// This field can be empty or omitted if the user is offline
	// or has the exact server turned off.
	Game string `json:"game,omitempty"`
}

// UserStats represents a user's statistics.
type UserStats struct {
	User  User                   `json:"user"`
	Stats map[string]interface{} `json:"stats"`
}

// UserAchievements represents a user's achievements.
type UserAchievements struct {
	User         User            `json:"user"`
	Achievements UserAchievement `json:"achievements"`
}

// UserAchievement represents a specific achievement of a user.
type UserAchievement struct {
	ID   int   `json:"id"`
	Time int64 `json:"time"`
}

// UserLeaderboards represents a user's positions on various leaderboards.
type UserLeaderboards struct {
	User         User              `json:"user"`
	Leaderboards []UserLeaderboard `json:"leaderboards"`
}

// UserLeaderboard represents a user's position on a specific leaderboard.
type UserLeaderboard struct {
	Type  string `json:"type"`
	Sort  string `json:"sort"`
	Place int    `json:"place"`
}

// UserMatches
type UserMatches struct {
	User    User         `json:"user"`
	Request MatchRequest `json:"request"`
	Matches []UserMatch  `json:"matches"`
}

// MatchRequest
type MatchRequest struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Size   int `json:"size"`
}

// Guild represents a guild, containing detailed information
// about the guild's properties, perks, and members.
type Guild struct {
	ID              int                  `json:"id"`
	Name            string               `json:"name"`
	Tag             string               `json:"tag"`
	Color           string               `json:"color"`
	Level           int                  `json:"level"`
	LevelPercentage float64              `json:"levelPercentage"`
	AvatarURL       string               `json:"avatar_url"`
	TotalExp        int                  `json:"totalExp"`
	TotalCoins      int                  `json:"totalCoins"`
	Created         int                  `json:"created"`
	WebInfo         string               `json:"web_info"`
	Perks           map[string]GuildPerk `json:"perks"`
	Members         []GuildMember        `json:"members"`
}

// GuildPerk represents a perk of the guild.
type GuildPerk struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// GuildMember represents a member of a guild.
type GuildMember struct {
	User       User   `json:"user"`
	Status     string `json:"status"`
	Joined     int    `json:"joined"`
	GuildCoins int    `json:"guildCoins"`
	GuildExp   int    `json:"guildExp"`
}

// GuildBrief represents a brief summary of the guild.
type GuildBrief struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Tag             string  `json:"tag"`
	Color           string  `json:"color"`
	Level           int     `json:"level"`
	LevelPercentage float64 `json:"levelPercentage"`
	AvatarURL       string  `json:"avatar_url"`
}

// Leaderboard represents a leaderboard.
type Leaderboard struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	MaxSize     int      `json:"max_size"`
	Sort        []string `json:"sort"`
}

// Online represents the total number of users currently online.
type Online struct {
	Total     int            `json:"total"`
	Separated map[string]int `json:"separated"`
}

// OnlineStreams represents live streams currently online.
type OnlineStreams struct {
	Title    string `json:"title"`
	Owner    string `json:"owner"`
	Viewers  int    `json:"viewers"`
	URL      string `json:"url"`
	Duration int    `json:"duration"`
	Platform string `json:"platform"`
	User     User   `json:"user"`
}

// OnlineStaff represents the online status of staff members.
type OnlineStaff struct {
	User
	Online OnlineSession `json:"online"`
}

// MatchBrief
type MatchBrief struct {
	ID       string   `json:"id"`
	Game     string   `json:"game"`
	Map      MatchMap `json:"map"`
	Date     int      `json:"date"`
	Duration int      `json:"duration"`
	Players  int      `json:"players"`
}

// MatchMap represents a match map.
type MatchMap struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Teams         int    `json:"teams"`
	PlayersInTeam int    `json:"playersInTeam"`
}

// UserMatch
type UserMatch struct {
	MatchBrief
	Win   bool `json:"win"`
	State int  `json:"state"`
}

// Locale represents localized game names, statistics, and ranks.
// https://vimeworld.github.io/api-docs/#apilocale_get
type Locale struct {
	Games     map[string]GameLocale             `json:"games"`
	GameStats map[string]map[string]interface{} `json:"game_stats"`
	Ranks     map[string]RankLocale             `json:"ranks"`
}

// GameLocale represents a localized name for a game.
type GameLocale struct {
	Name string `json:"name"`
}

// RankLocale represents a localized rank.
type RankLocale struct {
	Name     string   `json:"name"`
	Prefix   string   `json:"prefix"`
	Color    string   `json:"color"`
	Gradient []string `json:"gradient"`
}

// Game represents a game.
type Game struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// GlobalStats is a list of global statistics tracked for the game.
	GlobalStats []string `json:"global_stats"`

	// SeasonStats contains seasonal statistics tracked for the game.
	SeasonStats struct {
		Monthly []string `json:"monthly"`
		Manual  []string `json:"manual"`
	} `json:"season_stats"`
}

// Maps represents a collection of game maps.
type Maps map[string]map[string]Map

// Map represents a specific game map.
type Map struct {
	Name          string `json:"name"`
	Teams         int    `json:"teams"`
	PlayersInTeam int    `json:"playersInTeam"`
}

// Achievements represents a collection of achievements.
type Achievements map[string][]Achievement

// Achievement represents a single achievement that a user can earn.
type Achievement struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Reward      int      `json:"reward"`
	Description []string `json:"description"`
}

// Token represents information of a token.
// https://vimeworld.github.io/api-docs/#apimisc_token_get
type Token struct {
	Token string `json:"token"`
	Valid bool   `json:"valid"`
	Type  string `json:"type,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Owner *User  `json:"owner,omitempty"`
}
