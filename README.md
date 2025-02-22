[![Go Report Card](https://goreportcard.com/badge/github.com/0ffsideCompass/api-football-go-client)](https://github.com/0ffsideCompass/api-football-go-client)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://pkg.go.dev/badge/github.com/0ffsideCompass/api-football-go-client.svg)](https://pkg.go.dev/github.com/0ffsideCompass/api-football-go-client)

# API Football Go Client

A simple, intuitive, and comprehensive Go client library for interacting with the [API-Football](https://www.api-football.com/) service. This library allows you to easily retrieve data about leagues, teams, players, fixtures, transfers, and much more.

> **Note:** This client is designed with ease-of-use and flexibility in mind. It’s fully open source, well-documented, and built for developers who want to integrate API-Football data into their Go applications. Contributions, feedback, and improvements are welcome!

## Features

- **Clean & Simple API:** Access various endpoints (Coachs, Fixtures, Injuries, Leagues, Players, Standings, Teams, Transfers, Trophies, Venues, etc.) with intuitive methods.
- **Dynamic Query Parameters:** Pass custom filters and options via maps for flexible querying.
- **Customizable HTTP Client:** Inject your own `http.Client` for advanced configurations or testing purposes.
- **Robust Error Handling:** Clear error messages with wrapped errors to simplify debugging.
- **Comprehensive Documentation:** Inline comments and examples to help you get started quickly.

## Installation

Use `go get` to install the package:

```bash
go get github.com/0ffsideCompass/api-football-go-client
```

Or add it to your `go.mod` file:

```go
require github.com/0ffsideCompass/api-football-go-client vX.Y.Z
```

Replace `vX.Y.Z` with the current version.

## Quick Start

Here’s a simple example to get you started:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/0ffsideCompass/api-football-go-client/client"
)

func main() {
    // Replace with your API key
    apiKey := "YOUR_API_KEY_HERE"

    // Create a new client using the default domain
    cli, err := client.New(apiKey, &http.Client{Timeout: 10 * time.Second})
    if err != nil {
        log.Fatalf("Error creating client: %v", err)
    }

    // Example: Fetch leagues for the 2020 season in England
    params := map[string]interface{}{
        "season":  2020,
        "country": "England",
    }
    leagues, err := cli.Leagues(params)
    if err != nil {
        log.Fatalf("Error fetching leagues: %v", err)
    }

    fmt.Printf("Leagues: %+v\n", leagues)
}
```

## Endpoints Documentation

This client supports multiple endpoints. Below is an overview of some key methods:

### Coachs
- **Description:** Retrieves coach information.
- **Parameters:** One or more of `team`, `id`, or `search` (with `search` requiring at least 3 characters).

### Fixtures
- **Methods:** 
  - `Fixture` – Get fixtures with flexible filtering (by id, league, season, date, etc.).
  - `FixturesLineups` – Get lineups for a given fixture.
  - `FixturesEvents` – Retrieve events for a fixture.
  - `FixtureHeadToHead` – Get head-to-head stats between teams.
  - `FixtureByDateAndLeague` – Get fixtures by date range and league.
  - `FixtureStatistics` – Retrieve statistics for a fixture.
  - `FixturesPlayer` – Get player data for a fixture.

### Injuries
- **Description:** Retrieves injury reports.
- **Validation:** Ensures that if `league`, `team`, or `player` is provided, then `season` is required.

### Leagues
- **Description:** Retrieves league information with filters for `id`, `name`, `season`, `country`, etc.

### Players
- **Description:** Retrieve player details, seasons, squads, and top performance metrics.
- **Methods:** 
  - `PlayersSeasons`
  - `Players`
  - `PlayersSquads`
  - `PlayersTopScorers`
  - `PlayersTopAssists`
  - `PlayersTopYellowCards`
  - `PlayersTopRedCards`

### Standings
- **Description:** Get league standings.
- **Parameters:** Requires `season` and `league` (optionally `team`).

### Teams
- **Description:** Retrieves team data and statistics.
- **Methods:**
  - `Teams`
  - `TeamsStatistics`

### Transfers, Trophies, Venues
- **Description:** Access data related to player transfers, trophies, and venue information.

### Search
- **Description:** A generic search function for teams, leagues, or players.
- **Usage:** `Search(query string, t Type)` where `Type` is one of `Team`, `League`, or `Player`.

For detailed usage of each method, refer to the inline comments within the code.

## Contributing

Contributions are welcome! If you’d like to improve the client, please follow these guidelines:

- Fork the repository and create a new branch for your feature or bug fix.
- Ensure your code follows the existing style and documentation practices.
- Add tests for any new functionality.
- Open a pull request with a clear description of your changes.

## License

This project is licensed under the [MIT License](LICENSE).

## Support

If you encounter issues or have questions, please open an issue on [GitHub](https://github.com/0ffsideCompass/api-football-go-client/issues).

---

Happy coding and enjoy integrating API-Football data into your projects!