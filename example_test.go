package client_test

import (
	"fmt"
	"log"
	"net/http"
	"time"

	client "github.com/0ffsideCompass/api-football-go-client"
)

// The examples below are compiled but not executed, since they would require
// a real API key and network access.

func ExampleNew() {
	cli, err := client.New("YOUR_API_KEY", &http.Client{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}
	fmt.Println(cli.Domain)
}

func ExampleClient_Leagues() {
	cli, err := client.New("YOUR_API_KEY", &http.Client{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}

	leagues, err := cli.Leagues(map[string]any{
		"country": "England",
		"season":  2023,
	})
	if err != nil {
		log.Fatalf("fetching leagues: %v", err)
	}

	for i := range leagues.Response {
		fmt.Println(leagues.Response[i].League.Name)
	}
}

func ExampleClient_Fixture() {
	cli, err := client.New("YOUR_API_KEY", &http.Client{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}

	fixtures, err := cli.Fixture(map[string]any{
		"league": 39,
		"season": 2023,
		"date":   "2023-08-12",
	})
	if err != nil {
		log.Fatalf("fetching fixtures: %v", err)
	}

	fmt.Println(fixtures.Results)
}

func ExampleClient_FixtureByDateAndLeague() {
	cli, err := client.New("YOUR_API_KEY", &http.Client{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}

	from := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC)
	fixtures, err := cli.FixtureByDateAndLeague(39, 2023, from, to)
	if err != nil {
		log.Fatalf("fetching fixtures: %v", err)
	}

	fmt.Println(fixtures.Results)
}

func ExampleClient_Predictions() {
	cli, err := client.New("YOUR_API_KEY", &http.Client{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}

	predictions, err := cli.Predictions(map[string]any{"fixture": 198772})
	if err != nil {
		log.Fatalf("fetching predictions: %v", err)
	}

	for i := range predictions.Response {
		p := &predictions.Response[i].Predictions
		fmt.Printf("%s: %s\n", p.Winner.Name, p.Advice)
	}
}

func ExampleClient_Search() {
	cli, err := client.New("YOUR_API_KEY", &http.Client{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}

	// Search returns the raw JSON response body.
	body, err := cli.Search("Arsenal", client.Team)
	if err != nil {
		log.Fatalf("searching teams: %v", err)
	}

	fmt.Println(string(body))
}
