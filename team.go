package client

import (
	"encoding/json"
	"fmt"

	model "github.com/0ffsideCompass/api-football-adapter-model"
)

const (
	getTeamEndpoint      = "/team/get/%s/%s/%s"
	getTeamBasicEndpoint = "/team/get/%s"
)

// GetTeam sends a GET request to retrieve comprehensive team data including squad, venue, and fixtures.
// This function constructs a URL with the team ID, league ID, and season, sends a GET request to that URL,
// and parses the JSON response into a TeamResponse structure.
//
// Parameters:
//   - teamID: A string representing the unique identifier of the team
//   - leagueID: A string representing the unique identifier of the league
//   - season: A string representing the season (e.g., "2023")
//
// Returns:
//   - *TeamResponse: A pointer to the TeamResponse struct containing comprehensive team information including:
//   - Team basic information (name, logo, country, founded year)
//   - Squad with player details (name, position, number, nationality, etc.)
//   - Venue information specific to the league/season context
//   - Recent fixtures (last 5 completed matches with scores)
//   - Upcoming fixtures (next 5 scheduled matches)
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON.
func (c *Client) GetTeam(teamID, leagueID, season string) (*model.TeamResponse, error) {
	endpoint := fmt.Sprintf(getTeamEndpoint, teamID, leagueID, season)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error retrieving team data: %w", err)
	}

	var data model.TeamResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling team response data: %w", err)
	}

	return &data, nil
}

// GetTeamBasic sends a GET request to retrieve basic team data without league context.
// This function constructs a URL with only the team ID, sends a GET request to that URL,
// and parses the JSON response into a TeamResponse structure with basic team info and venue.
//
// Parameters:
//   - teamID: A string representing the unique identifier of the team
//
// Returns:
//   - *TeamResponse: A pointer to the TeamResponse struct containing basic team information including:
//     - Team basic information (name, logo, country, founded year)
//     - Venue information (team's home venue)
//     - Empty squad array (no league/season context)
//     - Empty fixtures arrays (no league/season context)
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON.
func (c *Client) GetTeamBasic(teamID string) (*model.TeamResponse, error) {
	endpoint := fmt.Sprintf(getTeamBasicEndpoint, teamID)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error retrieving basic team data: %w", err)
	}

	var data model.TeamResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling team response data: %w", err)
	}

	return &data, nil
}
