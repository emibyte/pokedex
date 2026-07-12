package player

import (
	"errors"
	"time"

	"github.com/emibyte/pokedex/internal/pokeapi"
)

type StatKind int

const (
	HP             StatKind = iota + 1 // "hp"
	Attack                             // "attack"
	Defense                            // "defense"
	SpecialAttack                      // "special-attack"
	SpecialDefense                     // "special-defense"
	Speed                              // "speed"
)

type CombatStat struct {
	Kind        StatKind
	BaseValue   int
	EffortValue int
}

type PartyPokemon struct {
	Name             string
	Nickname         string
	CaughtAt         time.Time
	Types            []string
	Level            int
	XP               int
	XPThresholdLevel int
	MaxHP            int
	CurrentHP        int
	BaseExperience   int
	CombatStats      []CombatStat
	GrowthRate       func(level int) (amount int)
}

func CreatePartyPokemon(pokemon pokeapi.Pokemon, species pokeapi.PokemonSpecies, nickname string) (PartyPokemon, error) {
	types := make([]string, len(pokemon.Types))
	for _, ty := range pokemon.Types {
		slot := ty.Slot
		typeName := ty.Type.Name

		idx := slot - 1
		types[idx] = typeName
	}

	combatStats := make([]CombatStat, len(pokemon.Stats))
	for _, st := range pokemon.Stats {
		kind, err := translateStatNameToKind(st.Stat.Name)
		if err != nil {
			return PartyPokemon{}, err
		}
		idx := translateStatKindToIdx(kind)
		combatStats[idx] = CombatStat{
			Kind:        kind,
			BaseValue:   st.BaseStat,
			EffortValue: st.Effort,
		}
	}

	growthRateFunc := getGrowthRateFunction(species.Name)

	p := PartyPokemon{
		Name:           pokemon.Name,
		Nickname:       nickname,
		CaughtAt:       time.Now(),
		Types:          types,
		Level:          1,
		XP:             0,
		BaseExperience: pokemon.BaseExperience,
		CombatStats:    combatStats,
		GrowthRate:     growthRateFunc,
	}

	p.MaxHP = p.GetCurrentStat(HP)
	p.CurrentHP = p.MaxHP

	p.XPThresholdLevel = p.GrowthRate(p.Level)

	return p, nil
}

func (p *PartyPokemon) GetCurrentStat(kind StatKind) int {
	idx := translateStatKindToIdx(kind)
	base := p.CombatStats[idx].BaseValue
	ev := p.CombatStats[idx].EffortValue

	return ((base + (ev / 4)) * p.Level / 50) + 5
}

func (p *PartyPokemon) HasFainted() bool {
	return p.CurrentHP < 1
}

func (p *PartyPokemon) GainXP(amount int) (levelsGained int) {
	p.XP += amount
	levelsGained = 0

	for p.XP > p.XPThresholdLevel {
		p.Level++
		levelsGained++

		p.XPThresholdLevel = p.GrowthRate(p.Level)

		oldMaxHP := p.MaxHP
		p.MaxHP = p.GetCurrentStat(HP)
		p.CurrentHP += p.MaxHP - oldMaxHP
	}

	return levelsGained
}

func translateStatNameToKind(name string) (StatKind, error) {
	switch name {
	case "hp":
		return HP, nil
	case "attack":
		return Attack, nil
	case "defense":
		return Defense, nil
	case "special-attack":
		return SpecialAttack, nil
	case "special-defense":
		return SpecialDefense, nil
	case "speed":
		return Speed, nil
	default:
		return 0, errors.New("Error: Stat name is not mapped")
	}
}

func translateStatKindToIdx(kind StatKind) int {
	return int(kind) - 1
}
