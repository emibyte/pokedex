package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + pokemonEndpoint + name
	data, found := c.cache.Get(url)

	if !found {
		resp, err := c.httpClient.Get(url)
		if err != nil {
			return Pokemon{}, fmt.Errorf("couldn't get pokemon data from pokeapi: %w", err)
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("couldn't read http body: %w", err)
		}

		if resp.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, data)
		}

		// NOTE: add to cache if no errors
		c.cache.Add(url, data)
	}

	var pokemon Pokemon
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("couldn't unmarshal data: %w", err)
	}

	return pokemon, nil
}
