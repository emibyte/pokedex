package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocation(locationName string) (LocationResponse, error) {
	url := baseUrl + locationAreaEndpoint + locationName
	data, found := c.cache.Get(url)

	if !found {
		resp, err := c.httpClient.Get(url)
		if err != nil {
			return LocationResponse{}, fmt.Errorf("couldn't get data from pokeapi: %w", err)
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationResponse{}, fmt.Errorf("couldn't read http body: %w", err)
		}

		if resp.StatusCode > 299 {
			return LocationResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, data)
		}

		// NOTE: add to cache if no errors
		c.cache.Add(url, data)
	}

	var location LocationResponse
	err := json.Unmarshal(data, &location)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("couldn't unmarshal data: %w", err)
	}

	return location, nil
}
