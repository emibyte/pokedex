package player

import "fmt"

const MaxPartySize = 6

type Party struct {
	Slots [6]PartyPokemon
	Count int
}

func (p *Party) isPartyFull() bool {
	return p.Count == MaxPartySize
}

func (p *Party) GetPartyPokemon(slot int) (PartyPokemon, error) {
	idx := slot - 1
	empty := PartyPokemon{}
	if idx < 0 {
		return empty, fmt.Errorf("cannot access a negative index")
	}
	if idx > p.Count-1 {
		return empty, fmt.Errorf("cannot access an index outside of the pokemon you possess")
	}

	return p.Slots[idx], nil
}

func (p *Party) AddToParty(pokemon PartyPokemon) bool {
	if p.isPartyFull() {
		return false
	}
	p.Slots[p.Count] = pokemon
	return true
}

func (p *Party) Remove(slot int) (PartyPokemon, error) {
	idx := slot - 1
	if idx > p.Count-1 {
		return PartyPokemon{}, fmt.Errorf("invalid idx: there's only %d Pokemon and the input was number %d", p.Count, idx+1)
	}
	if idx < 0 {
		return PartyPokemon{}, fmt.Errorf("invalid idx: negative index")
	}

	pokemon := p.Slots[idx]
	for i := idx; i < p.Count-1; i++ {
		p.Slots[i] = p.Slots[i+1]
	}

	p.Slots[p.Count-1] = PartyPokemon{}
	p.Count = p.Count - 1
	return pokemon, nil
}

func (p *Party) ListParty() {
	for i, pokemon := range p.Slots {
		displayName := pokemon.Name
		if pokemon.Nickname != "" {
			displayName = pokemon.Nickname
		}
		fmt.Printf("%d. %s\n", i, displayName)
		fmt.Printf("HP: %d/%d\n", pokemon.CurrentHP, pokemon.MaxHP)
	}
}
