package game

import (
	"fmt"
	"math/rand"
)

type PokemonStore struct {
	Pokemons []Pokemon
}

type Pokemon struct {
	ID      string
	Name    string
	Type    string
	ElmType string
	Health  int
	Moves   []Move
}

type Move struct {
	Name   string
	Type   string
	Damage int
}

func (ps *PokemonStore) GetPokemonById(pokemonId string) (*Pokemon, error) {
	for _, currPokemon := range ps.Pokemons {
		if currPokemon.ID == pokemonId {
			return &currPokemon, nil
		}
	}
	return nil, fmt.Errorf("no pokemon with id %s found in pokemonStore", pokemonId)
}

func (ps *PokemonStore) CatchPokemon(newPokemon Pokemon, name string) {
	newPokemon.Name = name
	newPokemon.ID = fmt.Sprintf("%d", rand.Int())
	ps.Pokemons = append(ps.Pokemons, newPokemon)
}

func (ps *PokemonStore) RemovePokemon(pokemonID string) error {
	newPokemon := []Pokemon{}
	removedPokemon := false
	for _, currPokemon := range ps.Pokemons {
		if currPokemon.ID != pokemonID {
			newPokemon = append(newPokemon, currPokemon)
		} else {
			if !removedPokemon {
				removedPokemon = true
			} else {
				return fmt.Errorf("2 pokemon of the id %s found in inventory", pokemonID)
			}
		}
	}
	if !removedPokemon {
		return fmt.Errorf("no item with id %s located in inventory", pokemonID)
	}
	ps.Pokemons = newPokemon
	return nil
}
