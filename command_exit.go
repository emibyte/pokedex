package main

import (
	"fmt"
	"os"
)

func commandExit(config *config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	savePokedex(config, config.repo)
	os.Exit(0)
	return nil
}
