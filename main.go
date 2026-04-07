package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/emibyte/pokedex/internal/pokeapi"
	"github.com/emibyte/pokedex/internal/pokepersistence"
)

func loadPokedex(repo pokepersistence.Repository) (*config, error) {
	// TODO: pull this out, what if we wanna load extra stuff (next/prevUrl or (in the future) current pokemon roster...)
	state, err := repo.Load()
	if err != nil {
		return &config{}, err
	}
	pokedex := make(map[string]pokeapi.Pokemon)
	json.Unmarshal(state, &pokedex)
	config := config{
		pokeapiClient: pokeapi.NewClient(time.Second * 5),
		pokedex:       pokedex,
		repo:          repo,
	}
	return &config, nil
}

func savePokedex(config *config, repo pokepersistence.Repository) error {
	state, err := json.Marshal(config.pokedex)
	if err != nil {
		return err
	}
	err = repo.Save(state)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	initCommandRegistry()

	// TODO: this should be in main not in startRepl
	repo, err := pokepersistence.InitFileRepository(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	config, err := loadPokedex(repo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer savePokedex(config, repo)
	startRepl(config)
}
