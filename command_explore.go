package main

import "fmt"

func commandExplore(config *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no area name was provided")
	}
	if len(args) > 1 {
		return fmt.Errorf("can only explore one area at a time!")
	}

	loc, err := config.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	if len(loc.PokemonEncounters) == 0 {

	}

	fmt.Println("Exploring ", args[0], "...")
	fmt.Println("Found Pokemon:")

	for _, encounter := range loc.PokemonEncounters {
		fmt.Println("- ", encounter.Pokemon.Name)
	}
	return nil
}
