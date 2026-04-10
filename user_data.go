package main

import (
	"time"

	"github.com/emibyte/pokedex/internal/pokeapi"
	"github.com/emibyte/pokedex/internal/pokepersistence"
)

func loadState(repo pokepersistence.Repository) (*config, error) {
	userData, err := pokepersistence.LoadUserData(repo)
	if err != nil {
		return &config{}, err
	}
	return mapUserDataToConfig(userData, repo), nil
}

func saveState(config *config, repo pokepersistence.Repository) error {
	userData := mapConfigToUserData(config)
	err := pokepersistence.SaveUserData(userData, repo)
	if err != nil {
		return err
	}
	return nil
}

func mapUserDataToConfig(userData *pokepersistence.UserData, repo pokepersistence.Repository) *config {
	config := config{
		pokeapiClient: pokeapi.NewClient(time.Second * 5),
		repo:          repo,
		pokedex:       userData.Pokedex,
		nextUrl:       userData.NextUrl,
		previousUrl:   userData.PreviousUrl,
	}

	return &config
}

func mapConfigToUserData(config *config) *pokepersistence.UserData {
	userData := pokepersistence.UserData{
		Pokedex:     config.pokedex,
		NextUrl:     config.nextUrl,
		PreviousUrl: config.previousUrl,
	}

	return &userData
}
