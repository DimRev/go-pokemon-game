package game

type GameStore struct {
	User         User
	ItemStore    ItemStore
	PokemonStore PokemonStore
}

func Test() {
	g := InitGame("Dima", 0)
	g.PokemonStore.CatchPokemon(PokemonMap["Charmander"], "My Charmander")
	pok := g.PokemonStore.Pokemons[0]
	pok.DisplayStats()
}

func InitGame(n string, s int) GameStore {
	return GameStore{
		User: User{
			Name:  n,
			Score: s,
		},
		ItemStore: ItemStore{
			Items: []Item{},
		},
		PokemonStore: PokemonStore{
			Pokemons: []Pokemon{},
		},
	}
}
