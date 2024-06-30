package game

import (
	"fmt"
)

type GameStore struct {
	User         User
	Inventory    Inventory
	PokemonStore PokemonStore
}

type User struct {
	Name  string
	Score int
}

type Inventory struct {
	Items []Item
}

type Item struct {
	ID          string
	Name        string
	Description string
}

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

func main() {
	g := InitGame("Dima", 0)
	pok := Pokemon{
		ID:      "ID",
		Name:    "Something",
		Type:    "J",
		ElmType: "S",
		Health:  100,
		Moves:   []Move{},
	}
	g.PokemonStore.CatchPokemon(pok)

}

func InitGame(n string, s int) GameStore {
	return GameStore{
		User: User{
			Name:  n,
			Score: s,
		},
		Inventory: Inventory{
			Items: []Item{},
		},
		PokemonStore: PokemonStore{
			Pokemons: []Pokemon{},
		},
	}
}

func (inv *Inventory) AddItemToInventory(newItem Item) {
	inv.Items = append(inv.Items, newItem)
}

func (inv *Inventory) RemoveItemFromInventory(itemID string) error {
	newItems := []Item{}
	removedItem := false
	for _, currItem := range inv.Items {
		if currItem.ID != itemID {
			newItems = append(newItems, currItem)
		} else {
			if !removedItem {
				removedItem = true
			} else {
				return fmt.Errorf("2 items of the id %s found in inventory", itemID)
			}
		}
	}
	if !removedItem {
		return fmt.Errorf("no item with id %s located in inventory", itemID)
	}
	inv.Items = newItems
	return nil
}

func (ps *PokemonStore) GetPokemonById(pokemonId string) (Pokemon, error) {
	for _, currPokemon := range ps.Pokemons {
		if currPokemon.ID == pokemonId {
			return currPokemon, nil
		}
	}
	return Pokemon{}, fmt.Errorf("no pokemon with id %s found in inventory", pokemonId)
}

func (ps *PokemonStore) CatchPokemon(newPokemon Pokemon) {
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

func (pok *Pokemon) DisplayStats() {
	fmt.Println("Name:   ", pok.Name)
	fmt.Println("Type:   ", pok.Type)
	fmt.Println("Health: ", pok.ElmType)
}
