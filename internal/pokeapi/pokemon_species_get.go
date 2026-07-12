package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

// TODO(emi): these are all really similar functions, should probably factor out the similar stuff
func (c *Client) GetPokemonSpecies(name string) (PokemonSpecies, error) {
	url := baseUrl + pokemonSpeciesEndpoint + name
	data, found := c.cache.Get(url)

	if !found {
		resp, err := c.httpClient.Get(url)
		if err != nil {
			return PokemonSpecies{}, fmt.Errorf("couldn't get pokemon species data from pokeapi: %w", err)
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokemonSpecies{}, fmt.Errorf("couldn't read http body: %w", err)
		}

		if resp.StatusCode > 299 {
			return PokemonSpecies{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, data)
		}

		// NOTE: add to cache if no errors
		c.cache.Add(url, data)
	}

	var species PokemonSpecies
	err := json.Unmarshal(data, &species)
	if err != nil {
		return PokemonSpecies{}, fmt.Errorf("couldn't unmarshal data: %w", err)
	}

	return species, nil
}
