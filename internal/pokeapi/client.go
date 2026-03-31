package pokeapi

import (
	"net/http"
	"time"

	"github.com/emibyte/pokedex/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2/"
const locationAreaEndpoint = "location-area/"
const pokemonEndpoint = "pokemon/"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(time.Second * 5),
	}
}
