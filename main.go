package main

import (
	"fmt"

	"github.com/emibyte/pokedex/internal/pokepersistence"
)

func main() {
	initCommandRegistry()

	repo, err := pokepersistence.InitFileRepository(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	config, err := loadState(repo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer saveState(config, repo)
	startRepl(config)
}
