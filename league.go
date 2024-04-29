package client

import (
	"encoding/json"
	"fmt"
)

const (
	addLeagueEndpoint = "/league/add"
	getLeagueEndpoint = "/league/get"
)

// AddFixture sends a POST request to add a new fixture based on the provided fixtureID.
// The function constructs a JSON payload with the fixtureID and posts it to the specified endpoint.
// It parses the response JSON into a GeneralFixtureData structure if successful.
//
// Parameters:
//   - fixtureID: A string representing the unique identifier of the fixture to be added.
//
// Returns:
//   - *GeneralFixtureData: A pointer to a GeneralFixtureData struct that contains detailed information about the fixture.
//   - error: An error object that captures issues during the POST request, response handling, or JSON parsing.
func (c *Client) AddLeague(leagueID, season string) (*LeagueAddResponse, error) {
	payload := LeagueRequest{
		LeagueID: leagueID,
		Season:   season,
	}

	responseData, err := c.post(addLeagueEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("error posting league: %w", err)
	}

	var data LeagueAddResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &data, nil
}

// GetFixture sends a GET request to retrieve details about a specific fixture by its ID.
// The function formats the endpoint URL with the fixture ID, sends the GET request,
// and parses the JSON response into a GeneralFixtureData structure.
//
// Parameters:
//   - id: A string representing the unique identifier of the fixture to retrieve.
//
// Returns:
//   - *GeneralFixtureData: A pointer to the GeneralFixtureData struct containing detailed information about the fixture.
//   - error: An error object that captures issues during the GET request, response handling, or JSON parsing.
func (c *Client) GetLeague(leagueID, season string) (*League, error) {
	payload := LeagueRequest{
		LeagueID: leagueID,
		Season:   season,
	}

	responseData, err := c.post(getLeagueEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("error retrieving league: %w", err)
	}

	var data League
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &data, nil
}
