package pokeapi

import (
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"
const locationAreaEndpoint = "location-area/"

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
