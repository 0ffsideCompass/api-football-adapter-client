# API Football Adapter Client

A Go client library for interacting with the api-football-adapter microservice. This client provides a simple interface to retrieve football data including teams, fixtures, and leagues.

## Installation

```bash
go get github.com/0ffsideCompass/api-football-adapter-client
```

## Quick Start

```go
import "github.com/0ffsideCompass/api-football-adapter-client"

// Initialize the client
client, err := client.New("https://your-api-url.com", "your-api-key")
if err != nil {
    log.Fatal(err)
}

// Get a team with full details
team, err := client.GetTeam("33", "39", "2023") // Arsenal, Premier League, 2023 season
if err != nil {
    log.Fatal(err)
}

// Get basic team information
basicTeam, err := client.GetTeamBasic("33")
if err != nil {
    log.Fatal(err)
}
```

## Features

### Team Operations
- **GetTeam(teamID, leagueID, season)** - Get comprehensive team data including squad, venue, and fixtures
- **GetTeamBasic(teamID)** - Get basic team information without league context

### Fixture Operations
- **AddFixture(fixtureID)** - Add a new fixture to the system
- **GetFixture(id)** - Retrieve details about a specific fixture
- **GetFixtureByDateAndLeague(date, league)** - Get fixtures by date and league

### League Operations
- **AddLeague(leagueID, season)** - Add a new league for a specific season
- **GetLeague(leagueID, season)** - Get league information for a specific season

## Authentication

The client requires an API key for authentication. The API key is passed as a header in all requests:

```go
client, err := client.New("https://api.example.com", "your-api-key-here")
```

## Error Handling

All methods return detailed error information. Common error scenarios include:
- Invalid API key or URL during client initialization
- Network connectivity issues
- Invalid request parameters
- JSON unmarshalling errors

## Dependencies

- [api-football-adapter-model](https://github.com/0ffsideCompass/api-football-adapter-model) - Data models
- [api-football-go-client](https://github.com/0ffsideCompass/api-football-go-client) - Additional API client models

## Contributing

This project is part of the OffsideCompass football data ecosystem. For questions or contributions, please refer to the main project documentation.