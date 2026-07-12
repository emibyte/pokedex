package player

type CaughtPokemonResult int

const (
	AddedToParty CaughtPokemonResult = iota
	AddedToPc
)

type Player struct {
	Party Party
	PC    []PartyPokemon
	Name  string
	Money int
}

func CreatePlayer(name string) *Player {
	// NOTE(emi): Party is already initialized since its a fixed size array
	player := Player{
		PC:    make([]PartyPokemon, 0, 256),
		Name:  name,
		Money: 500,
	}
	return &player
}

func (p *Player) AddCaughtPokemon(pokemon PartyPokemon) CaughtPokemonResult {
	added := p.Party.AddToParty(pokemon)
	if !added {
		p.PC = append(p.PC, pokemon)
		return AddedToPc
	}
	return AddedToParty
}

func (p *Player) MovePokemonToPC(slot int) {
}
