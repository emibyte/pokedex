package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

func (c *Client) GetTypeMatchupStats() ([]TypeDetail, error) {
	url := baseUrl + pokemonTypeEndpoint
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return []TypeDetail{}, fmt.Errorf("couldnt get type listing from pokeapi: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []TypeDetail{}, fmt.Errorf("couldn't read http body: %w", err)
	}

	if resp.StatusCode > 299 {
		return []TypeDetail{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, data)
	}

	var list TypeList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return []TypeDetail{}, fmt.Errorf("couldn't unmarshal data: %w", err)
	}

	var wg sync.WaitGroup
	results := make([]TypeDetail, len(list.Results))
	for i, result := range list.Results {
		wg.Add(1)
		go func(idx int, url string) {
			defer wg.Done()
			typeResp, err := c.httpClient.Get(url)
			if err != nil {
				return
			}
			defer typeResp.Body.Close()

			if typeResp.StatusCode > 299 {
				return
			}

			data, err := io.ReadAll(typeResp.Body)
			_ = json.Unmarshal(data, &results[idx])

		}(i, result.URL)
	}
	wg.Wait()

	return results, nil
}
