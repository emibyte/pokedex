package main

import (
	"fmt"
)

func commandMapb(config *config, _ ...string) error {
	if config.previousUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locs, err := config.pokeapiClient.ListLocations(config.previousUrl)
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
