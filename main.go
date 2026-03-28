package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			cleanedInput := cleanInput(text)
			if len(cleanedInput) == 0 {
				continue
			}
			fmt.Printf("Your command was: %s\n", cleanedInput[0])
		}
	}
}
