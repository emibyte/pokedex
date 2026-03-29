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
