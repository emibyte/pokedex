package pokepersistence

import (
	"encoding/json"

	"github.com/emibyte/pokedex/internal/pokeapi"
)

// NOTE: struct that contains everything in the config we intend to persist somewhere
// this and config needs to match up with each other
// (haven't thought of an easier way to do this yet sadly, maybe some reflection stuff?)
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
		// NOTE: if there is no file (content empty), return the 0 value of UserData
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
