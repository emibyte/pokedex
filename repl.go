package main

import "strings"

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	// NOTE: Fields does the thing that i just did manually before (also catches tabs and stuff)
	words := strings.Fields(lower)
	return words
}
