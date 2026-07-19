// Package client provides a Go client for the API-Football v3 service from
// api-sports.io, covering every endpoint of the API: leagues, teams, fixtures,
// standings, players, transfers, injuries, predictions, odds, and more.
//
// # Creating a client
//
// A client needs an API key and an HTTP client. Any type implementing the
// HttpClient interface (a single Do method) works, so a plain *http.Client is
// enough:
//
//	cli, err := client.New(apiKey, &http.Client{Timeout: 10 * time.Second})
//	if err != nil {
//		log.Fatal(err)
//	}
//
// By default requests go to the direct API at https://v3.football.api-sports.io/,
// authenticated with the x-apisports-key header. To use the RapidAPI gateway
// instead, pass its base URL to NewWithDomain; the client sends the RapidAPI
// headers as well, so both providers work without further configuration.
//
// # Calling endpoints
//
// Most methods take a map of query parameters mirroring the API's own
// parameter names, and return a typed response model from the models package:
//
//	fixtures, err := cli.Fixture(map[string]any{
//		"league": 39,
//		"season": 2023,
//		"date":   "2023-08-12",
//	})
//
// Each method's documentation lists the parameters the endpoint accepts,
// including which ones are required. Required parameters are validated before
// the request is sent. Endpoints that take no parameters (for example
// Timezone and LeaguesSeasons) take no arguments.
//
// # Error handling
//
// Methods return wrapped errors for failed requests (non-2xx responses,
// with the response body included in the message) and for undecodable
// responses.
//
// Be aware that API-Football reports some failures — invalid parameters,
// exceeded quotas — inside the body of a 200 OK response. The client does not
// currently turn those into errors; check the Errors field on the response
// model to distinguish an empty result from a rejected request.
//
// # Retries and rate limits
//
// The client ships without retry logic by design: inject a retrying HTTP
// client (such as hashicorp/go-retryablehttp's StandardClient) if you need
// it. Note that API-Football enforces per-minute and per-day quotas and every
// request counts against them, so avoid blindly retrying 429 responses within
// the same rate window.
package client
