package battlesim

import (
	"fmt"

	"github.com/emibyte/pokedex/internal/pokeapi"
)

type cacheEntry struct {
	doubleDamageTo []string
	halfDamageTo   []string
	noDamageTo     []string
}

type typeMatchupCache = map[string]cacheEntry

var typeMatchups typeMatchupCache = make(map[string]cacheEntry)

func buildCache(client *pokeapi.Client) error {
	typeDetails, err := client.GetTypeMatchupStats()
	if err != nil {
		return fmt.Errorf("couldnt fetch type details for cache: %w")
	}

	for _, typeDetail := range typeDetails {
		entry := cacheEntry{
			doubleDamageTo: make([]string, len(typeDetail.DamageRelations.DoubleDamageTo)),
			halfDamageTo:   make([]string, len(typeDetail.DamageRelations.HalfDamageTo)),
			noDamageTo:     make([]string, len(typeDetail.DamageRelations.NoDamageTo)),
		}

		extractTypeNames(typeDetail.DamageRelations.DoubleDamageTo, entry.doubleDamageTo)
		extractTypeNames(typeDetail.DamageRelations.HalfDamageTo, entry.halfDamageTo)
		extractTypeNames(typeDetail.DamageRelations.NoDamageTo, entry.noDamageTo)

		typeMatchups[typeDetail.Name] = entry
	}

	return nil
}

func extractTypeNames(typeInfos []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}, types []string) {
	for i, typeInfo := range typeInfos {
		types[i] = typeInfo.Name
	}
}

func getMultiplier(attackerType, defenderType string) float64 {
	eff := typeMatchups[attackerType]
	for _, ty := range eff.doubleDamageTo {
		if ty == defenderType {
			return 2.0
		}
	}
	for _, ty := range eff.halfDamageTo {
		if ty == defenderType {
			return 0.5
		}
	}
	for _, ty := range eff.noDamageTo {
		if ty == defenderType {
			return 0.0
		}
	}
	return 1.0
}
