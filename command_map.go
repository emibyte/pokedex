package main

import (
	"fmt"
)

func commandMap(config *config) error {
	locs, err := config.pokeapiClient.ListLocations(config.nextUrl)
	if err != nil {
		return err
	}

	config.nextUrl = locs.Next
	config.previousUrl = locs.Previous
	for _, res := range locs.Results {
		fmt.Println(res.Name)
	}
	return nil
}
