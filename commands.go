package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
}

func registerCommand(registry commandRegistry, name, description string, callback func() error) error {
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	var sb strings.Builder
	beginning := "Welcome to the Pokedex!\nUsage:\n\n"
	sb.WriteString(beginning)
	for _, cmd := range registry.commands {
		fmt.Fprintf(&sb, "%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println(sb.String())
	return nil
}
