package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func (c *Client) ListLocations(pageUrl *string) (LocationsResponse, error) {
	url := baseUrl + locationAreaEndpoint
	if pageUrl != nil {
		url = *pageUrl
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationsResponse{}, fmt.Errorf("couldn't get data from pokeapi: %w", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, data)
	}

	if err != nil {
		return LocationsResponse{}, fmt.Errorf("couldn't read http body: %w", err)
	}

	var locations LocationsResponse
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsResponse{}, fmt.Errorf("couldn't unmarshal data: %w", err)
	}

	return locations, nil
}
