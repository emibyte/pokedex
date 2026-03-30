package main

import (
	"fmt"
	"strings"
)

func commandHelp(_ *config, _ ...string) error {
	var sb strings.Builder
	beginning := "Welcome to the Pokedex!\nUsage:\n\n"
	sb.WriteString(beginning)
	for _, cmd := range registry.commands {
		fmt.Fprintf(&sb, "%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println(sb.String())
	return nil
}
