package main

import "fmt"

func commandPokedex(config *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
