package pokeapi

type LocationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`     // can be null so pointer necessary
	Previous *string `json:"previous"` // can be null so pointer necessary
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationResponse struct {
	// NOTE: the actual response is way bigger but since we only care about the encounters we only map this
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
