package game

var PokemonMap = map[string]Pokemon{
	"Charmander": {
		ID:      "placeholder-id",
		Name:    "placeholder-name",
		Type:    "Charmander",
		ElmType: "Fire",
		Health:  100,
		Moves: []Move{
			{Name: "Flamethrower", Type: "Fire", Damage: 50},
			{Name: "Scratch", Type: "Normal", Damage: 10},
		},
	},
	"Pikachu": {
		ID:      "placeholder-id",
		Name:    "placeholder-name",
		Type:    "Pikachu",
		ElmType: "Electric",
		Health:  100,
		Moves: []Move{
			{Name: "Thunderbolt", Type: "Electric", Damage: 50},
			{Name: "Quick Attack", Type: "Normal", Damage: 10},
		},
	},
	"Bulbasaur": {
		ID:      "placeholder-id",
		Name:    "placeholder-name",
		Type:    "Bulbasaur",
		ElmType: "Grass/Poison",
		Health:  100,
		Moves: []Move{
			{Name: "Vine Whip", Type: "Grass", Damage: 45},
			{Name: "Tackle", Type: "Normal", Damage: 10},
		},
	},
	"Squirtle": {
		ID:      "placeholder-id",
		Name:    "placeholder-name",
		Type:    "Water",
		ElmType: "Water",
		Health:  100,
		Moves: []Move{
			{Name: "Water Gun", Type: "Water", Damage: 40},
			{Name: "Tackle", Type: "Normal", Damage: 10},
		},
	},
}
