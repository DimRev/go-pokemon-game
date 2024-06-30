package game

import "log"

type GameStore struct {
	User         User
	ItemStore    ItemStore
	PokemonStore PokemonStore
}

type User struct {
	Name  string
	Score int
}

func Test() {
	player := InitUser("Dima", 0)
	computer := InitOpponent("COM", 0)
	player.PokemonStore.CatchPokemon(PokemonMap["Charmander"], "Player Char")
	computer.PokemonStore.CatchPokemon(PokemonMap["Pikachu"], "Computer Pika")

	PlayerPok1, err := player.PokemonStore.GetPokemonByIdx(1)
	if err != nil {
		log.Fatal(err)
	}
	ComputerPok1, err := computer.PokemonStore.GetPokemonByIdx(1)
	if err != nil {
		log.Fatal(err)
	}
	player.PokemonStore.DisplayPokemons()

	PlayerPok1.UseMoveOn(1, ComputerPok1)
	ComputerPok1.UseMoveOn(1, PlayerPok1)
	PlayerPok1.UseMoveOn(1, ComputerPok1)
	ComputerPok1.UseMoveOn(1, PlayerPok1)
}

func InitUser(n string, s int) GameStore {
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

func InitOpponent(n string, s int) GameStore {
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
