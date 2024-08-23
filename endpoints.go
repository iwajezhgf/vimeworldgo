package vimeworldgo

// Known VimeWorld API endpoints.
var (
	EndpointAPI         = "https://api.vimeworld.com/"
	EndpointUser        = EndpointAPI + "user/"
	EndpointGuild       = EndpointAPI + "guild/"
	EndpointLeaderboard = EndpointAPI + "leaderboard/"
	EndpointOnline      = EndpointAPI + "online"
	EndpointLocale      = EndpointAPI + "locale/"
	EndpointMisc        = EndpointAPI + "misc/"

	EndpointLeaderboardList  = EndpointLeaderboard + "list"
	EndpointOnlineStreams    = EndpointOnline + "/streams"
	EndpointOnlineStaff      = EndpointOnline + "/staff"
	EndpointMiscGames        = EndpointMisc + "games"
	EndpointMiscMaps         = EndpointMisc + "maps"
	EndpointMiscAchievements = EndpointMisc + "achievements"

	EndpointUserNames        = func(names string) string { return EndpointUser + "name/" + names }
	EndpointUserIds          = func(ids string) string { return EndpointUser + ids }
	EndpointUserFriends      = func(id string) string { return EndpointUser + id + "/friends" }
	EndpointUserSession      = func(id string) string { return EndpointUser + id + "/session" }
	EndpointUserStats        = func(id, games string) string { return EndpointUser + id + "/stats" + games }
	EndpointUserAchievements = func(id string) string { return EndpointUser + id + "/achievements" }
	EndpointUserLeaderboards = func(id string) string { return EndpointUser + id + "/leaderboards" }
	EndpointUserMatches      = func(id, query string) string { return EndpointUser + id + "/matches" + query }
	EndpointGuildSearch      = func(query string) string { return EndpointGuild + "search" + query }
	EndpointGuildGet         = func(query string) string { return EndpointGuild + "get" + query }
	EndpointLocaleName       = func(name, query string) string { return EndpointLocale + name + query }
	EndpointMiscToken        = func(token string) string { return EndpointMisc + "token/" + token }
)
