package battlesim

import (
	"fmt"

	"github.com/emibyte/pokedex/internal/player"
	"github.com/emibyte/pokedex/internal/pokeapi"
)

// NOTE(emi): everything here is really simplified, this doesn't really work like a real turn based
//            battle between two pokemon (we're just emulating sth that appears like it :c)

type BattleResult struct {
	Won          bool
	LeveledUp    bool
	LevelsGained int
}

func RunSimulation(client *pokeapi.Client, playerPokemon *player.PartyPokemon, enemy pokeapi.Pokemon, enemySpecies pokeapi.PokemonSpecies) (BattleResult, error) {
	enemyPokemon, err := player.CreatePartyPokemon(enemy, enemySpecies, "")
	if err != nil {
		return BattleResult{}, fmt.Errorf("error creating enemy pokemon: %w", err)
	}

	err = stageBattle(client)
	if err != nil {
		return BattleResult{}, fmt.Errorf("error staging battle: %w", err)
	}

	return runBattle(playerPokemon, &enemyPokemon), nil
}

func stageBattle(client *pokeapi.Client) error {
	if len(typeMatchups) == 0 {
		err := buildCache(client)
		if err != nil {
			return fmt.Errorf("error initializing type matchup cache: %w", err)
		}
	}
	return nil
}

func runBattle(playerPKMN, enemyPKMN *player.PartyPokemon) BattleResult {
	result := BattleResult{}

	playerMult := getTypeMatchupDmgMultiplier(playerPKMN, enemyPKMN)
	enemyMult := getTypeMatchupDmgMultiplier(enemyPKMN, playerPKMN)

	playerGoesFirst := playerPKMN.GetCurrentStat(player.Speed) >= enemyPKMN.GetCurrentStat(player.Speed)

	for playerPKMN.CurrentHP > 0 && enemyPKMN.CurrentHP > 0 {
		if playerGoesFirst {
			// Player attacks, then check if enemy fainted
			// If enemy still alive, enemy attacks
			executeAttack(playerPKMN, enemyPKMN, playerMult)
			if enemyPKMN.HasFainted() {
				break
			}
			executeAttack(enemyPKMN, playerPKMN, enemyMult)
		} else {
			// Enemy attacks, then check if player fainted
			// If player still alive, player attacks
			executeAttack(enemyPKMN, playerPKMN, enemyMult)
			if playerPKMN.HasFainted() {
				break
			}
			executeAttack(playerPKMN, enemyPKMN, playerMult)
		}
	}

	result.Won = !playerPKMN.HasFainted()
	if result.Won {
		// NOTE(emi): some level scaling on it to reward beating a pokemon thats higher level than urs
		xpToAward := (enemyPKMN.BaseExperience * enemyPKMN.Level) / (playerPKMN.Level)
		gained := playerPKMN.GainXP(xpToAward)
		result.LeveledUp = gained > 0
		result.LevelsGained = gained
	}

	return result
}

func getTypeMatchupDmgMultiplier(source, target *player.PartyPokemon) float64 {
	totalMultiplier := 1.0
	attackerPrimaryType := source.Types[0]

	for _, defenderType := range source.Types {
		mult := getMultiplier(attackerPrimaryType, defenderType)
		totalMultiplier *= mult
	}

	return totalMultiplier
}

func executeAttack(attacker, defender *player.PartyPokemon, multiplier float64) {
	attackStat := attacker.GetCurrentStat(player.Attack)
	defenseStat := attacker.GetCurrentStat(player.Defense)

	movePower := 40
	baseDamage := float64(((2 * attacker.Level / 5) + 2) * movePower * attackStat / defenseStat)
	finalDamage := (baseDamage / 50) + 2

	totalDamage := multiplier * finalDamage
	defender.CurrentHP -= int(totalDamage)
	if defender.CurrentHP < 0 {
		defender.CurrentHP = 0
	}
}
