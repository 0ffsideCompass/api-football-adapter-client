package client

import (
	"encoding/json"
	"fmt"

	model "github.com/0ffsideCompass/api-football-adapter-model"
)

const (
	getPlayerEndpoint      = "/player/get/%s/%s/%s"
	getPlayerBasicEndpoint = "/player/get/%s"
)

// GetPlayer sends a GET request to retrieve comprehensive player data including statistics, transfers, and performance data.
// This function constructs a URL with the player ID, league ID, and season, sends a GET request to that URL,
// and parses the JSON response into a PlayerResponse structure.
//
// Parameters:
//   - playerID: A string representing the unique identifier of the player
//   - leagueID: A string representing the unique identifier of the league
//   - season: A string representing the season (e.g., "2023")
//
// Returns:
//   - *PlayerResponse: A pointer to the PlayerResponse struct containing comprehensive player information including:
//     - Player basic information (name, age, nationality, physical attributes)
//     - Current season statistics for the specified league
//     - Career statistics across all leagues and seasons
//     - Transfer history (when available)
//     - Injury history (when available)
//     - Performance data and season highlights
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON.
func (c *Client) GetPlayer(playerID, leagueID, season string) (*model.PlayerResponse, error) {
	endpoint := fmt.Sprintf(getPlayerEndpoint, playerID, leagueID, season)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error retrieving player data: %w", err)
	}

	var data model.PlayerResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling player response data: %w", err)
	}

	return &data, nil
}

// GetPlayerBasic sends a GET request to retrieve basic player data without league context.
// This function constructs a URL with only the player ID, sends a GET request to that URL,
// and parses the JSON response into a PlayerResponse structure with basic player info and current season stats.
//
// Parameters:
//   - playerID: A string representing the unique identifier of the player
//
// Returns:
//   - *PlayerResponse: A pointer to the PlayerResponse struct containing basic player information including:
//     - Player basic information (name, age, nationality, physical attributes)
//     - Current season statistics (across all competitions)
//     - Career statistics for all available seasons
//     - Empty transfer and injury history arrays (not fetched in basic mode)
//     - Basic performance data without detailed match analysis
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON.
func (c *Client) GetPlayerBasic(playerID string) (*model.PlayerResponse, error) {
	endpoint := fmt.Sprintf(getPlayerBasicEndpoint, playerID)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error retrieving basic player data: %w", err)
	}

	var data model.PlayerResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling player response data: %w", err)
	}

	return &data, nil
}