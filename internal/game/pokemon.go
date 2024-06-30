package game

import (
	"errors"
	"fmt"
)

func (pok *Pokemon) DisplayStats() {
	fmt.Println("Name:   ", pok.Name)
	fmt.Println("Type:   ", pok.Type)
	fmt.Println("Health: ", pok.ElmType)
	pok.ShowMoves()
}

func (pok *Pokemon) ShowMoves() {
	if len(pok.Moves) > 0 {
		fmt.Println("Moves:  ")
		for idx, move := range pok.Moves {
			fmt.Printf("\t%d) %s\t    %d(%s)\n", idx+1, move.Name, move.Damage, move.Type)
		}
	}
}

func (pok *Pokemon) GetMoveByIdx(moveIdx int) (Move, error) {
	if moveIdx < 0 || moveIdx >= len(pok.Moves) {
		return Move{}, errors.New("invalid move index")
	}
	return pok.Moves[moveIdx], nil
}
