[![CI](https://github.com/0ffsideCompass/api-football-go-client/actions/workflows/ci.yml/badge.svg)](https://github.com/0ffsideCompass/api-football-go-client/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/0ffsideCompass/api-football-go-client/branch/main/graph/badge.svg)](https://codecov.io/gh/0ffsideCompass/api-football-go-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/0ffsideCompass/api-football-go-client)](https://goreportcard.com/report/0ffsideCompass/api-football-go-client)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://pkg.go.dev/badge/github.com/0ffsideCompass/api-football-go-client.svg)](https://pkg.go.dev/github.com/0ffsideCompass/api-football-go-client)

# API Football Go Client

A simple, intuitive, and comprehensive Go client library for interacting with the [API-Football](https://www.api-football.com/) service. This library allows you to easily retrieve data about leagues, teams, players, fixtures, transfers, and much more.

> **Note:** This client is designed with ease-of-use and flexibility in mind. It’s fully open source, well-documented, and built for developers who want to integrate API-Football data into their Go applications. Contributions, feedback, and improvements are welcome!

## Features

- **Complete Endpoint Coverage:** Every API-Football v3 endpoint is supported — Coachs, Countries, Fixtures, Injuries, Leagues, Odds (pre-match and in-play), Players, Predictions, Sidelined, Standings, Teams, Timezone, Transfers, Trophies, and Venues.
- **Dynamic Query Parameters:** Pass custom filters and options via maps for flexible querying.
- **Customizable HTTP Client:** Inject your own `http.Client` for advanced configurations or testing purposes.
- **Wrapped Errors:** Non-2xx responses and JSON decoding failures are returned as wrapped errors to simplify debugging (see [Error Handling](#error-handling) for a caveat about API-level errors).
- **Works with Both API Providers:** Compatible with the direct api-sports.io API and the RapidAPI gateway (see [Authentication](#authentication)).

## Installation

Use `go get` to install the package:

```bash
go get github.com/0ffsideCompass/api-football-go-client
```

## Quick Start

Here’s a simple example to get you started:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    client "github.com/0ffsideCompass/api-football-go-client"
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
    params := map[string]any{
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

## Authentication

The client sends both authentication headers on every request, so it works with either provider without configuration:

- **Direct API** (`https://v3.football.api-sports.io/`, the default domain): authenticated via the `x-apisports-key` header, using the API key from your [api-football.com](https://www.api-football.com/) dashboard.
- **RapidAPI**: use `client.NewWithDomain` with the RapidAPI base URL (`https://api-football-v1.p.rapidapi.com/v3/`) and your RapidAPI key, which is sent via the `X-RapidAPI-Key` header.

## Error Handling

Methods return an error for failed requests (non-2xx status codes, including the response body for context) and for responses that cannot be decoded.

**Caveat:** API-Football reports some failures — invalid parameters, exceeded rate limits — inside the response body of a `200 OK` response. The client does not currently inspect that field, so such calls return an empty result rather than an error. Check the `Errors` field on the response model if you need to distinguish "no data" from "bad request".

## Endpoints Documentation

This client supports multiple endpoints. Below is an overview of some key methods:

### Coachs
- **Description:** Retrieves coach information.
- **Parameters:** One or more of `team`, `id`, or `search` (with `search` requiring at least 3 characters).

### Fixtures
- **Methods:** 
  - `Fixture` – Get fixtures with flexible filtering (by id, league, season, date, etc.).
  - `FixturesRounds` – Get the rounds of a league (optionally with dates).
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
- **Methods:** `Leagues`, `LeaguesSeasons` (list of available 4-digit seasons).

### Predictions
- **Description:** Retrieves predictions for a fixture, including both teams' form, comparative statistics, and head-to-head history.
- **Parameters:** Requires `fixture`.

### Odds
- **Methods:**
  - `Odds` – Pre-match odds (filter by `fixture`, `league`, `season`, `date`, `bookmaker`, `bet`, `page`).
  - `OddsMapping` – Fixture IDs with available pre-match odds.
  - `OddsBookmakers` / `OddsBets` – Available bookmakers and bet types.
  - `OddsLive` – In-play odds for fixtures in progress.
  - `OddsLiveBets` – Available bet types for in-play odds.

### Countries & Timezone
- **Methods:** `Countries` (filter by `name`, `code`, `search`), `Timezone` (no parameters).

### Players
- **Description:** Retrieve player details, seasons, squads, and top performance metrics.
- **Methods:** 
  - `PlayersSeasons`
  - `Players`
  - `PlayersProfiles`
  - `PlayersSquads`
  - `PlayersTeams`
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
  - `TeamsSeasons`
  - `TeamsCountries`

### Transfers, Trophies, Venues, Sidelined
- **Description:** Access data related to player transfers, trophies, venue information, and sidelined players/coaches.

### Search
- **Description:** A generic search function for teams, leagues, or players.
- **Usage:** `Search(query string, t Type)` where `Type` is one of `Team`, `League`, or `Player`.
- **Note:** Returns the raw JSON response body (`[]byte`) rather than a typed model. Queries must be at least 3 characters (4 for players).

For detailed usage of each method, refer to the [package documentation](https://pkg.go.dev/github.com/0ffsideCompass/api-football-go-client) or the inline comments within the code.

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