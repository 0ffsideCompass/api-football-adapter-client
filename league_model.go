package client

// LeagueRequest represents the request to retrieve a league's data.
type LeagueRequest struct {
	LeagueID string `json:"league_id""`
	Season   string `json:"season""`
}

// LeagueAddResponse represents the response to adding a league's data.
type LeagueAddResponse struct {
	Message string `json:"message"`
}

// League represents a football league including its basic details and the teams and standings within it.
type League struct {
	ID           string      `json:"id" bson:"id"`
	Name         string      `json:"name" bson:"name"`
	Season       string      `json:"season" bson:"season"`
	SeasonNumber int         `json:"season_number" bson:"season_number"`
	Country      string      `json:"country" bson:"country"`
	Flag         string      `json:"flag" bson:"flag"`
	Finished     bool        `json:"finished" bson:"finished"`
	Standings    []Standings `json:"standings" bson:"standings"`
	TeamsPath    []TeamPath  `json:"teams_path" bson:"teams_path"`
}

// TeamPath represents the path of a team through the league, detailing each round and fixture.
type TeamPath struct {
	TeamID        string         `json:"team_id" bson:"team_id"`
	TeamName      string         `json:"team_name" bson:"team_name"`
	TeamLogo      string         `json:"team_logo" bson:"team_logo"`
	RoundFixtures []RoundFixture `json:"round_fixtures" bson:"round_fixtures"`
}

// RoundFixture provides the details of a single fixture in a league round for a specific team.
type RoundFixture struct {
	Round            string `json:"round" bson:"round"`
	RoundNum         int    `json:"round_number" bson:"round_number"`
	FixtureID        string `json:"fixture_id" bson:"fixture_id"`
	HomeGame         bool   `json:"home_game" bson:"home_game"`
	AgainstTeam      string `json:"against_team" bson:"against_team"`
	AgainstTeamID    string `json:"against_team_id" bson:"against_team_id"`
	ResultForTeam    string `json:"result_for_team" bson:"result_for_team"`
	Points           int    `json:"points" bson:"points"`
	Goals            int    `json:"goals" bson:"goals"`
	GoalsAgainst     int    `json:"goals_against" bson:"goals_against"`
	TotalGoal        int    `json:"total_goal" bson:"total_goal"`
	TotalGoalAgainst int    `json:"total_goal_against" bson:"total_goal_against"`
}

// Standings details the current standings of a team within its league.
type Standings struct {
	Rank             int    `json:"rank" bson:"rank"`
	Team             string `json:"team" bson:"team"`
	TeamID           int    `json:"team_id" bson:"team_id"`
	TeamLogo         string `json:"team_logo" bson:"team_logo"`
	Points           int    `json:"points" bson:"points"`
	GoalsDiff        int    `json:"goal_diff" bson:"goals_diff"`
	Form             string `json:"form" bson:"form"`
	Played           int    `json:"played" bson:"played"`
	Wins             int    `json:"wins" bson:"wins"`
	Draws            int    `json:"draws" bson:"draws"`
	Losses           int    `json:"losses" bson:"losses"`
	GoalsFor         int    `json:"goals_for" bson:"goals_for"`
	GoalsAgainst     int    `json:"goals_against" bson:"goals_against"`
	HomePlayed       int    `json:"home_played" bson:"home_played"`
	HomeWins         int    `json:"home_wins" bson:"home_wins"`
	HomeDraws        int    `json:"home_draws" bson:"home_draws"`
	HomeLosses       int    `json:"home_losses" bson:"home_losses"`
	HomeGoalsFor     int    `json:"home_goals_for" bson:"home_goals_for"`
	HomeGoalsAgainst int    `json:"home_goals_against" bson:"home_goals_against"`
	AwayPlayed       int    `json:"away_played" bson:"away_played"`
	AwayWins         int    `json:"away_wins" bson:"away_wins"`
	AwayDraws        int    `json:"away_draws" bson:"away_draws"`
	AwayLosses       int    `json:"away_losses" bson:"away_losses"`
	AwayGoalsFor     int    `json:"away_goals_for" bson:"away_goals_for"`
	AwayGoalsAgainst int    `json:"away_goals_against" bson:"away_goals_against"`
}
