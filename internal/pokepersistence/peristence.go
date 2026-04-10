package pokepersistence

import (
	"encoding/json"

	"github.com/emibyte/pokedex/internal/pokeapi"
)

// NOTE: struct that contains everything in the config we intend to persist somewhere
type UserData struct {
	Pokedex     map[string]pokeapi.Pokemon
	NextUrl     *string
	PreviousUrl *string
}

func LoadUserData(repo Repository) (*UserData, error) {
	state, err := repo.Load()
	if err != nil {
		return &UserData{}, err
	}

	if len(state) == 0 {
		return &UserData{
			Pokedex:     make(map[string]pokeapi.Pokemon),
			NextUrl:     nil,
			PreviousUrl: nil,
		}, nil
	}
	var data UserData
	err = json.Unmarshal(state, &data)
	if err != nil {
		return &UserData{}, err
	}
	return &data, nil
}

func SaveUserData(userData *UserData, repo Repository) error {
	state, err := json.Marshal(userData)
	if err != nil {
		return err
	}
	err = repo.Save(state)
	if err != nil {
		return err
	}
	return nil
}
