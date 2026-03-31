package main

import (
	"fmt"

	"github.com/emibyte/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
	nextUrl       *string
	previousUrl   *string
}

type commandRegistry struct {
	commands map[string]cliCommand
}

var registry = commandRegistry{
	commands: make(map[string]cliCommand),
}

func initCommandRegistry() {
	registerCommand(registry, "exit", "Exit the Pokedex", commandExit)
	registerCommand(registry, "help", "Displays a help message", commandHelp)
	registerCommand(registry, "map", "Get the next page of locations", commandMap)
	registerCommand(registry, "mapb", "Get the previous page of locations", commandMapb)
	registerCommand(registry, "explore", "Get the pokemon for a location area", commandExplore)
	registerCommand(registry, "catch", "Attempt to catch a pokemon", commandCatch)
}

func registerCommand(registry commandRegistry, name, description string, callback func(*config, ...string) error) error {
	cmd := cliCommand{
		name:        name,
		description: description,
		callback:    callback,
	}

	if _, ok := registry.commands[name]; !ok {
		registry.commands[name] = cmd
		return nil
	}
	return fmt.Errorf("command %s has already been added to the registry", name)
}
