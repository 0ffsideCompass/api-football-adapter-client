package client

import (
	"time"
)

// GetFixturesByDateRequest represents the request body for fetching fixtures by date.
type GetFixturesByDateAndLeagueRequest struct {
	Date   string `json:"date"`
	League string `json:"league"`
}

// FixtureRequest represents the request body fora specific fixture.
type FixtureRequest struct {
	FixtureID string `json:"fixture_id"`
}

// GeneralFixtureData represents the overall data structure for storing information
// about a single game fixture in a MongoDB database. It includes identifiers, standings data,
// detailed match information, and team performance statistics.
type GeneralFixtureData struct {
	StandingsData StandingsData  `json:"standings_data" bson:"standings"`
	FixtureID     string         `json:"fixture_id" bson:"fixture_id"`
	FixtureData   FixtureData    `json:"fixture_data" bson:"current_data"`
	HomeTeamStats TeamStatistics `json:"home_team_stats" bson:"home_form_data"`
	AwayTeamStats TeamStatistics `json:"away_team_stats" bson:"away_form_data"`
}

// FixtureData holds specific details about a match including the participating teams, venue,
// scores, and other relevant details.
type FixtureData struct {
	Referee           string    `json:"referee" bson:"referee"`
	Timezone          string    `json:"timezone" bson:"timezone"`
	Date              time.Time `json:"date" bson:"date"`
	Venue             string    `json:"venue" bson:"venue"`
	VanueCity         string    `json:"venue_city" bson:"venue_city"`
	GameStatus        string    `json:"game_status" bson:"game_status"`
	GameTime          int       `json:"game_time" bson:"game_time"`
	LeagueName        string    `json:"league_name" bson:"league_name"`
	LeagueCountry     string    `json:"league_country" bson:"league_country"`
	LeagueRound       string    `json:"league_round" bson:"league_round"`
	HomeTeam          string    `json:"home_team" bson:"home_team"`
	AwayTeam          string    `json:"away_team" bson:"away_team"`
	HomeTeamLogo      string    `json:"home_team_logo" bson:"home_team_logo"`
	AwayTeamLogo      string    `json:"away_team_logo" bson:"away_team_logo"`
	HomeTeamID        int       `json:"home_team_id" bson:"home_team_id"`
	AwayTeamID        int       `json:"away_team_id" bson:"away_team_id"`
	Winner            string    `json:"winner" bson:"winner"`
	GoalsHome         int       `json:"goals_home" bson:"goals_home"`
	GoalsAway         int       `json:"goals_away" bson:"goals_away"`
	ScoreHalfTimeHome int       `json:"score_halftime_home" bson:"score_halftime_home"`
	ScoreHalfTimeAway int       `json:"score_halftime_away" bson:"score_halftime_away"`
	ScoreFullTimeHome int       `json:"score_fulltime_home" bson:"score_fulltime_home"`
	ScoreFullTimeAway int       `json:"score_fulltime_away" bson:"score_fulltime_away"`
	ScoreExtraTime    int       `json:"score_extra_time" bson:"score_extra_time"`
	ScorePenatyHome   int       `json:"score_penalty_home" bson:"score_penalty_home"`
	ScorePenatyAway   int       `json:"score_penalty_away" bson:"score_penalty_away"`
	Events            []Event   `json:"events" bson:"events"`
	Finished          bool      `json:"finished" bson:"finished"`
	UpdateAt          time.Time `json:"update_at" bson:"update_at"`
}

// Event represents a significant occurrence during a fixture such as a goal, card, or substitution.
type Event struct {
	TimeElapsed int    `json:"time_elapsed" bson:"time_elapsed"`
	Team        string `json:"team" bson:"team"`
	Player      string `json:"player" bson:"player"`
	Assist      string `json:"assist" bson:"assist"`
	Type        string `json:"type" bson:"type"`
	Detail      string `json:"detail" bson:"detail"`
	Comments    string `json:"comments" bson:"comments"`
}

// StandingsData contains information about the current standings in a league.
type StandingsData struct {
	LeagueName string         `json:"league_name" bson:"name"`
	Standings  []TeamStanding `json:"standings" bson:"standings"`
}

// TeamStanding details the position and performance metrics of a team within its league.
type TeamStanding struct {
	Rank        int        `json:"rank" bson:"rank"`
	TeamName    string     `json:"team_name" bson:"team"`
	Points      int        `json:"points" bson:"points"`
	GoalsDiff   int        `json:"goals_diff" bson:"goalsdiff"`
	Group       string     `json:"group" bson:"group"`
	Form        string     `json:"form" bson:"form"`
	Status      string     `json:"status" bson:"status"`
	Description string     `json:"description" bson:"description"`
	All         PlayedData `json:"all" bson:"all"`
	Home        PlayedData `json:"home" bson:"home"`
	Away        PlayedData `json:"away" bson:"away"`
}

// PlayedData summarizes the performance of a team over a set of matches.
type PlayedData struct {
	Played int       `json:"played" bson:"played"`
	Win    int       `json:"win" bson:"win"`
	Draw   int       `json:"draw" bson:"draw"`
	Lose   int       `json:"lose" bson:"lose"`
	Goals  GoalsData `json:"goals" bson:"goals"`
}

// GoalsData holds the number of goals scored and conceded by a team.
type GoalsData struct {
	For     int `json:"for" bson:"for"`
	Against int `json:"against" bson:"against"`
}

// TeamStatistics contains detailed performance metrics for a team in a league.
type TeamStatistics struct {
	TeamName                       string   `json:"team_name" bson:"team"`
	Form                           string   `json:"form" bson:"form"`
	PlayedHome                     int      `json:"played_home" bson:"played_home"`
	PlayedAway                     int      `json:"played_away" bson:"played_away"`
	Total                          int      `json:"total" bson:"total"`
	WinsHome                       int      `json:"wins_home" bson:"win_home"`
	WinsAway                       int      `json:"wins_away" bson:"win_away"`
	WinsTotal                      int      `json:"wins_total" bson:"win_total"`
	DrawsHome                      int      `json:"draws_home" bson:"draw_home"`
	DrawsAway                      int      `json:"draws_away" bson:"draw_away"`
	DrawsTotal                     int      `json:"draws_total" bson:"draw_total"`
	LosesHome                      int      `json:"loses_home" bson:"lose_home"`
	LosesAway                      int      `json:"loses_away" bson:"lose_away"`
	LosesTotal                     int      `json:"loses_total" bson:"lose_total"`
	GoalsTotal                     int      `json:"goals_total" bson:"goals_total"`
	GoalsHome                      int      `json:"goals_home" bson:"goals_home"`
	GoalsAway                      int      `json:"goals_away" bson:"goals_away"`
	GoalAvgTotal                   string   `json:"goal_avg_total" bson:"goal_avg_total"`
	GoalAvgHome                    string   `json:"goal_avg_home" bson:"goal_avg_home"`
	GoalAvgAway                    string   `json:"goal_avg_away" bson:"goal_avg_away"`
	Goals0To15                     int      `json:"goals_0_to_15" bson:"goals_0_to_15"`
	Goals0To15Percentage           string   `json:"goals_0_to_15_percentage" bson:"goals_0_to_15_percentage"`
	Goals16To30                    int      `json:"goals_16_to_30" bson:"goals_16_to_30"`
	Goals16To30Percentage          string   `json:"goals_16_to_30_percentage" bson:"goals_16_to_30_percentage"`
	Goals31To45                    int      `json:"goals_31_to_45" bson:"goals_31_to_45"`
	Goals31To45Percentage          string   `json:"goals_31_to_45_percentage" bson:"goals_31_to_45_percentage"`
	Goals46To60                    int      `json:"goals_46_to_60" bson:"goals_46_to_60"`
	Goals46To60Percentage          string   `json:"goals_46_to_60_percentage" bson:"goals_46_to_60_percentage"`
	Goals61To75                    int      `json:"goals_61_to_75" bson:"goals_61_to_75"`
	Goals61To75Percentage          string   `json:"goals_61_to_75_percentage" bson:"goals_61_to_75_percentage"`
	Goals76To90                    int      `json:"goals_76_to_90" bson:"goals_76_to_90"`
	Goals76To90Percentage          string   `json:"goals_76_to_90_percentage" bson:"goals_76_to_90_percentage"`
	Goals91to105                   int      `json:"goals_91_to_105" bson:"goals_91_to_105"`
	Goals91to105Percentage         string   `json:"goals_91_to_105_percentage" bson:"goals_91_to_105_percentage"`
	Goals106To120                  int      `json:"goals_106_to_120" bson:"goals_106_to_120"`
	Goals106To120Percentage        string   `json:"goals_106_to_120_percentage" bson:"goals_106_to_120_percentage"`
	AgainstGoalTotal               int      `json:"against_goal_total" bson:"against_goal_total"`
	AgainstGoalHome                int      `json:"against_goal_home" bson:"against_goal_home"`
	AgainstGoalAway                int      `json:"against_goal_away" bson:"against_goal_away"`
	AgainstGoalAvgTotal            string   `json:"against_goal_avg_total" bson:"against_goal_avg_total"`
	AgainstGoalAvgHome             string   `json:"against_goal_avg_home" bson:"against_goal_avg_home"`
	AgainstGoalAvgAway             string   `json:"against_goal_avg_away" bson:"against_goal_avg_away"`
	AgainstGoals0To15              int      `json:"against_goals_0_to_15" bson:"against_goals_0_to_15"`
	AgainstGoals0To15Percentage    string   `json:"against_goals_0_to_15_percentage" bson:"against_goals_0_to_15_percentage"`
	AgainstGoals16To30             int      `json:"against_goals_16_to_30" bson:"against_goals_16_to_30"`
	AgainstGoals16To30Percentage   string   `json:"against_goals_16_to_30_percentage" bson:"against_goals_16_to_30_percentage"`
	AgainstGoals31To45             int      `json:"against_goals_31_to_45" bson:"against_goals_31_to_45"`
	AgainstGoals31To45Percentage   string   `json:"against_goals_31_to_45_percentage" bson:"against_goals_31_to_45_percentage"`
	AgainstGoals46To60             int      `json:"against_goals_46_to_60" bson:"against_goals_46_to_60"`
	AgainstGoals46To60Percentage   string   `json:"against_goals_46_to_60_percentage" bson:"against_goals_46_to_60_percentage"`
	AgainstGoals61To75             int      `json:"against_goals_61_to_75" bson:"against_goals_61_to_75"`
	AgainstGoals61To75Percentage   string   `json:"against_goals_61_to_75_percentage" bson:"against_goals_61_to_75_percentage"`
	AgainstGoals76To90             int      `json:"against_goals_76_to_90" bson:"against_goals_76_to_90"`
	AgainstGoals76To90Percentage   string   `json:"against_goals_76_to_90_percentage" bson:"against_goals_76_to_90_percentage"`
	AgainstGoals91to105            int      `json:"against_goals_91_to_105" bson:"against_goals_91_to_105"`
	AgainstGoals91to105Percentage  string   `json:"against_goals_91_to_105_percentage" bson:"against_goals_91_to_105_percentage"`
	AgainstGoals106To120           int      `json:"against_goals_106_to_120" bson:"against_goals_106_to_120"`
	AgainstGoals106To120Percentage string   `json:"against_goals_106_to_120_percentage" bson:"against_goals_106_to_120_percentage"`
	BiggestSteakWins               int      `json:"biggest_steak_wins" bson:"biggest_steak_wins"`
	BiggestSteakDraws              int      `json:"biggest_steak_draws" bson:"biggest_steak_draws"`
	BiggestSteakLoses              int      `json:"biggest_steak_loses" bson:"biggest_steak_loses"`
	BiggestWinsHome                string   `json:"biggest_wins_home" bson:"biggest_wins_home"`
	BiggestWinsAway                string   `json:"biggest_wins_away" bson:"biggest_wins_away"`
	BiggestLosesHome               string   `json:"biggest_loses_home" bson:"biggest_loses_home"`
	BiggestLosesAway               string   `json:"biggest_loses_away" bson:"biggest_loses_away"`
	BiggestGoalsForHome            int      `json:"biggest_goals_for_home" bson:"biggest_goals_for_home"`
	BiggestGoalsForAway            int      `json:"biggest_goals_for_away" bson:"biggest_goals_for_away"`
	BiggestGoalsAgainstHome        int      `json:"biggest_goals_against_home" bson:"biggest_goals_against_home"`
	BiggestGoalsAgainstAway        int      `json:"biggest_goals_against_away" bson:"biggest_goals_against_away"`
	CleanSheetsHome                int      `json:"clean_sheets_home" bson:"clean_sheets_home"`
	CleanSheetsAway                int      `json:"clean_sheets_away" bson:"clean_sheets_away"`
	CleanSheetsTotal               int      `json:"clean_sheets_total" bson:"clean_sheets_total"`
	FailedToScoreHome              int      `json:"failed_to_score_home" bson:"failed_to_score_home"`
	FailedToScoreAway              int      `json:"failed_to_score_away" bson:"failed_to_score_away"`
	FailedToScoreTotal             int      `json:"failed_to_score_total" bson:"failed_to_score_total"`
	PenaltyScoredTotal             int      `json:"penalty_scored_total" bson:"penalty_scored_total"`
	PenaltyScoredPercentage        string   `json:"penalty_scored_percentage" bson:"penalty_scored_percentage"`
	PenaltyMissedTotal             int      `json:"penalty_missed_total" bson:"penalty_missed_total"`
	PenaltyMissedPercentage        string   `json:"penalty_missed_percentage" bson:"penalty_missed_percentage"`
	PenaltyTotal                   int      `json:"penalty_total" bson:"penalty_total"`
	Lineups                        []Lineup `json:"lineups" bson:"lineups"`
	CardsYellow0to15Total          int      `json:"cards_yellow_0_to_15_total" bson:"cards_yellow_0_15_total"`
	CardsYellow0to15Percentage     string   `json:"cards_yellow_0_to_15_percentage" bson:"cards_yellow_0_15_percentage"`
	CardsYellow16to30Total         int      `json:"cards_yellow_16_to_30_total" bson:"cards_yellow_16_30_total"`
	CardsYellow16to30Percentage    string   `json:"cards_yellow_16_to_30_percentage" bson:"cards_yellow_16_30_percentage"`
	CardsYellow31to45Total         int      `json:"cards_yellow_31_to_45_total" bson:"cards_yellow_31_45_total"`
	CardsYellow31to45Percentage    string   `json:"cards_yellow_31_to_45_percentage" bson:"cards_yellow_31_45_percentage"`
	CardsYellow46to60Total         int      `json:"cards_yellow_46_to_60_total" bson:"cards_yellow_46_60_total"`
	CardsYellow46to60Percentage    string   `json:"cards_yellow_46_to_60_percentage" bson:"cards_yellow_46_60_percentage"`
	CardsYellow61to75Total         int      `json:"cards_yellow_61_to_75_total" bson:"cards_yellow_61_75_total"`
	CardsYellow61to75Percentage    string   `json:"cards_yellow_61_to_75_percentage" bson:"cards_yellow_61_75_percentage"`
	CardsYellow76to90Total         int      `json:"cards_yellow_76_to_90_total" bson:"cards_yellow_76_90_total"`
	CardsYellow76to90Percentage    string   `json:"cards_yellow_76_to_90_percentage" bson:"cards_yellow_76_90_percentage"`
	CardsYellow91to105Total        int      `json:"cards_yellow_91_to_105_total" bson:"cards_yellow_91_105_total"`
	CardsYellow91to105Percentage   string   `json:"cards_yellow_91_to_105_percentage" bson:"cards_yellow_91_105_percentage"`
	CardsYellow106to120Total       int      `json:"cards_yellow_106_to_120_total" bson:"cards_yellow_106_120_total"`
	CardsYellow106to120Percentage  string   `json:"cards_yellow_106_to_120_percentage" bson:"cards_yellow_106_120_percentage"`
	CardsRed0to15Total             int      `json:"cards_red_0_to_15_total" bson:"cards_red_0_15_total"`
	CardsRed0to15Percentage        string   `json:"cards_red_0_to_15_percentage" bson:"cards_red_0_15_percentage"`
	CardsRed16to30Total            int      `json:"cards_red_16_to_30_total" bson:"cards_red_16_30_total"`
	CardsRed16to30Percentage       string   `json:"cards_red_16_to_30_percentage" bson:"cards_red_16_30_percentage"`
	CardsRed31to45Total            int      `json:"cards_red_31_to_45_total" bson:"cards_red_31_45_total"`
	CardsRed31to45Percentage       string   `json:"cards_red_31_to_45_percentage" bson:"cards_red_31_45_percentage"`
	CardsRed46to60Total            int      `json:"cards_red_46_to_60_total" bson:"cards_red_46_60_total"`
	CardsRed46to60Percentage       string   `json:"cards_red_46_to_60_percentage" bson:"cards_red_46_60_percentage"`
	CardsRed61to75Total            int      `json:"cards_red_61_to_75_total" bson:"cards_red_61_75_total"`
	CardsRed61to75Percentage       string   `json:"cards_red_61_to_75_percentage" bson:"cards_red_61_75_percentage"`
	CardsRed76to90Total            int      `json:"cards_red_76_to_90_total" bson:"cards_red_76_90_total"`
	CardsRed76to90Percentage       string   `json:"cards_red_76_to_90_percentage" bson:"cards_red_76_90_percentage"`
	CardsRed91to105Total           int      `json:"cards_red_91_to_105_total" bson:"cards_red_91_105_total"`
	CardsRed91to105Percentage      string   `json:"cards_red_91_to_105_percentage" bson:"cards_red_91_105_percentage"`
	CardsRed106to120Total          int      `json:"cards_red_106_to_120_total" bson:"cards_red_106_120_total"`
	CardsRed106to120Percentage     string   `json:"cards_red_106_to_120_percentage" bson:"cards_red_106_120_percentage"`
}

// Lineup represents the formation and number of matches played with that formation.
type Lineup struct {
	Formation string `json:"formation" bson:"formation"`
	Played    int    `json:"played" bson:"played"`
}
