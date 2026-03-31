package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(config *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name was provided")
	}
	if len(args) > 1 {
		return fmt.Errorf("can only catch one pokemon at a time!")
	}

	pokemon, err := config.pokeapiClient.GetPokemon(strings.ToLower(args[0]))
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at", pokemon.Name+"...")

	// NOTE: the only important thing is that pokemon with higher baseExperience are harder
	//       to catch than pokemon with smaller baseExperience values
	//       this pattern leads to about a 1 in 2 chance for catching a pokemon with 50 base exp
	//       and a 1 in 5 for catching one with 200 base exp
	difficulty := max(pokemon.BaseExperience/50, 1)
	roll := rand.Intn(difficulty + 1)
	if roll == 0 {
		fmt.Println(pokemon.Name, "was caught!")
		config.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name, "was not caught, try again!")
	}

	return nil
}
