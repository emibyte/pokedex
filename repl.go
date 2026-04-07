package main

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	// NOTE: Fields does the thing that i just did manually before (also catches tabs and stuff)
	words := strings.Fields(lower)
	return words
}

func startRepl(config *config) {
	// TODO: implement this myself way too big of a dependency just for command history tbh
	rl, err := readline.New("Pokedex > ")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println(err)
			return
		}

		words := cleanInput(line)
		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := words[1:]
		cmd, exists := registry.commands[command]
		if exists {
			err := cmd.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
