package main

import (
	"fmt"
	"strings"
)

func commandInspect(config *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name was provided")
	}
	if len(args) > 1 {
		return fmt.Errorf("can only inspect one pokemon at a time!")
	}
	if pokemon, ok := config.pokedex[strings.ToLower(args[0])]; ok {
		fmt.Printf("Name:   %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, ty := range pokemon.Types {
			fmt.Printf("  -%s\n", ty.Type.Name)
		}
		return nil
	}
	fmt.Println("you have not caught that pokemon!")
	return nil
}
