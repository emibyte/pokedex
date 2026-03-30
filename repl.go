package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/emibyte/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	// NOTE: Fields does the thing that i just did manually before (also catches tabs and stuff)
	words := strings.Fields(lower)
	return words
}

func startRepl() {
	initCommandRegistry()
	config := config{
		pokeapiClient: pokeapi.NewClient(time.Second * 5),
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := words[1:]
		cmd, exists := registry.commands[command]
		if exists {
			err := cmd.callback(&config, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
