package client

import (
	"encoding/json"
	"fmt"

	model "github.com/0ffsideCompass/api-football-adapter-model"
	"github.com/0ffsideCompass/api-football-go-client/models"
)

const (
	addFixtureEndpoint                = "/fixture/add"
	getFixtureEndpoint                = "/fixture/get/%s"
	getFixtureByDateAndLeagueEndpoint = "/fixture/get/bydateandleague"
	getFixturesEndpoint               = "/fixtures"
)

// This function constructs a JSON payload that includes the fixture ID, sends it to the server using a POST request,
// and parses the JSON response into a GeneralFixtureData structure.
//
// Parameters:
//   - fixtureID: A string representing the unique identifier of the fixture to be added.
//
// Returns:
//   - *GeneralFixtureData: A pointer to the GeneralFixtureData struct containing detailed information about the fixture.
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON.
func (c *Client) AddFixture(fixtureID string) (*model.GeneralFixtureData, error) {
	payload := model.FixtureRequest{
		FixtureID: fixtureID,
	}

	responseData, err := c.post(addFixtureEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("error posting fixture: %w", err)
	}

	var data model.GeneralFixtureData
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &data, nil
}

// GetFixture sends a GET request to retrieve details about a specific fixture by its ID.
// This function constructs a URL with the fixture ID, sends a GET request to that URL,
// and parses the JSON response into a GeneralFixtureData structure.
//
// Parameters:
//   - id: A string representing the unique identifier of the fixture to retrieve.
//
// Returns:
//   - *GeneralFixtureData: A pointer to the GeneralFixtureData struct containing detailed information about the fixture.
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON.
func (c *Client) GetFixture(id string) (*model.GeneralFixtureData, error) {
	endpoint := fmt.Sprintf(getFixtureEndpoint, id)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error retrieving fixture: %w", err)
	}

	var data model.GeneralFixtureData
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &data, nil
}

func (c *Client) GetFixtureByDateAndLeague(date, league string) (*models.FixturesResponse, error) {
	var payload = model.GetFixturesByDateAndLeagueRequest{
		Date:   date,
		League: league,
	}

	responseData, err := c.post(getFixtureByDateAndLeagueEndpoint, payload)
	if err != nil {
		return nil, fmt.Errorf("error retrieving fixture: %w", err)
	}

	var data models.FixturesResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &data, nil
}

// GetFixtures retrieves fixtures based on flexible query parameters.
// This is a generic method that accepts any combination of parameters supported by API-Sports.io.
// It passes the parameters to the adapter service which forwards them to the Football API.
//
// Supported parameters include:
//   - id: (integer) The ID of a specific fixture
//   - ids: (string) Multiple fixture IDs in format "id-id-id" (max 20)
//   - live: (string) Get live fixtures - "all" for all leagues or "id-id" for specific leagues
//   - date: (string) Date in format "YYYY-MM-DD"
//   - league: (integer) The league ID
//   - season: (integer) The season year in 4-digit format (YYYY)
//   - team: (integer) The team ID
//   - last: (integer) Last N fixtures (max 2 digits)
//   - next: (integer) Next N fixtures (max 2 digits)
//   - from: (string) Start date for range in format "YYYY-MM-DD"
//   - to: (string) End date for range in format "YYYY-MM-DD"
//   - round: (string) The round of the fixture
//   - status: (string) Fixture status like "NS", "FT", "NS-PST-FT" (combined)
//   - venue: (integer) The venue ID
//   - timezone: (string) Timezone for fixture times
//
// Parameters:
//   - params: A map of query parameters to pass to the API
//
// Returns:
//   - *models.FixturesResponse: Response containing fixture data from the API
//   - error: An error object that reports issues in the request or response handling
//
// Example usage:
//
//	// Get fixture by ID
//	fixtures, err := client.GetFixtures(map[string]interface{}{"id": "215662"})
//
//	// Get today's fixtures
//	fixtures, err := client.GetFixtures(map[string]interface{}{"date": "2024-01-15"})
//
//	// Get all live fixtures
//	fixtures, err := client.GetFixtures(map[string]interface{}{"live": "all"})
//
//	// Get fixtures for a specific team
//	fixtures, err := client.GetFixtures(map[string]interface{}{
//		"team": 33,
//		"season": 2024,
//		"last": 5,
//	})
func (c *Client) GetFixtures(params map[string]interface{}) (*models.FixturesResponse, error) {
	responseData, err := c.post(getFixturesEndpoint, params)
	if err != nil {
		return nil, fmt.Errorf("error retrieving fixtures: %w", err)
	}

	var data models.FixturesResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling fixtures response: %w", err)
	}

	return &data, nil
}
