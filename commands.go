package main

import (
	"fmt"

	"github.com/emibyte/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

type config struct {
	pokeapiClient pokeapi.Client
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
	registerCommand(registry, "map", "Displays the name of 20 location areas, each subsequent call displays the next 20", commandMap)
}

func registerCommand(registry commandRegistry, name, description string, callback func(*config) error) error {
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
