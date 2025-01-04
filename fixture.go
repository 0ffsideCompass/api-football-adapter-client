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
