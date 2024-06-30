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

func (ps *PokemonStore) DisplayPokemons() {
	pokemonStr := ""
	if len(ps.Pokemons) == 0 {
		fmt.Println("No pokemons")
	}
	for idx, pokemon := range ps.Pokemons {
		if idx == len(ps.Pokemons)-1 {
			pokemonStr += fmt.Sprintf("[%d]%s", idx+1, pokemon.Name)
		} else if idx == 0 {
			pokemonStr += fmt.Sprintf("[%d]%s|", idx+1, pokemon.Name)
		} else {
			if (idx+1)%3 == 0 {
				pokemonStr += fmt.Sprintf("[%d]%s\n", idx+1, pokemon.Name)
			} else {
				pokemonStr += fmt.Sprintf("[%d]%s|", idx+1, pokemon.Name)
			}
		}

	}
	fmt.Println(pokemonStr)
}

func (ps *PokemonStore) GetPokemonByIdx(pokemonIdx int) (*Pokemon, error) {
	for idx, currPokemon := range ps.Pokemons {
		if idx == pokemonIdx-1 {
			return &currPokemon, nil
		}
	}
	return nil, fmt.Errorf("no pokemon on %d position found in pokemonStore", pokemonIdx)
}

func (ps *PokemonStore) CatchPokemon(newPokemon Pokemon, name string) {
	newPokemon.Name = name
	newPokemon.ID = fmt.Sprintf("%d", rand.Int())
	ps.Pokemons = append(ps.Pokemons, newPokemon)
}

func (ps *PokemonStore) RemovePokemon(pokemonIdx int) error {
	newPokemon := []Pokemon{}
	removedPokemon := false
	for idx, currPokemon := range ps.Pokemons {
		if idx != pokemonIdx {
			newPokemon = append(newPokemon, currPokemon)
		} else {
			removedPokemon = true
		}
	}
	if !removedPokemon {
		return fmt.Errorf("no pokemons found in %d position", pokemonIdx)
	}
	ps.Pokemons = newPokemon
	return nil
}
