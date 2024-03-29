package main

import (
	"fmt"
	"net/http"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func main() {
	//apiKey := "xxxxxxxxx"

	//client, err := client.New(apiKey, &http.Client{})
	//if err != nil {
	//		panic(err)
	//}

	cl, err := client.NewWithDomain(
		"4393126f4amshe521f523d3a9035p1acfd0jsnc2a0f68cf6b5",
		"https://api-football-v1.p.rapidapi.com/v3/",
		&http.Client{},
	)

	if err != nil {
		panic(err)
	}

	x, _ := cl.Search("Arsenl", client.Team)

	fmt.Println(string(x))

	// Get all fixtures for a given date and league
	//_, err = client.FixtureByDateAndLeague(
	//	2,
	//	2020,
	//	time.Date(2020, 12, 1, 0, 0, 0, 0, time.UTC),
	//	time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
	//)

	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//fmt.Println(fixtures)

	// Get head to head for two teams
	//headToHead, err := client.FixtureHeadToHead(33, 34)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//fmt.Println(headToHead)

	/*queryParameters := make(map[string]interface{})
	queryParameters["live"] = "all"

	fixture, err := client.Fixture(queryParameters)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(fixture)

	*/
	/*queryParameters := make(map[string]interface{})
	queryParameters["city"] = "Madrid"
	vanues, err := client.Venues(queryParameters)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(vanues)
	*/

	//queryParameters := make(map[string]interface{})
	//queryParameters["season"] = 2023
	//queryParameters["league"] = 39
	//vanues, err := cl.Standings(queryParameters)
	//if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(vanues)
}
