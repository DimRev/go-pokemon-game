package game

import (
	"errors"
	"fmt"
)

func (pok *Pokemon) DisplayStats() {
	fmt.Println("Name:   ", pok.Name)
	fmt.Println("Type:   ", pok.Type)
	fmt.Println("Health: ", pok.Health)
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
	if moveIdx-1 < 0 || moveIdx-1 >= len(pok.Moves) {
		return Move{}, errors.New("invalid move index")
	}
	return pok.Moves[moveIdx-1], nil
}

func (pok *Pokemon) UseMoveOn(moveIdx int, pok2 *Pokemon) (isAlive bool, err error) {
	mv, err := pok.GetMoveByIdx(moveIdx)
	if err != nil {
		return false, err
	}
	fmt.Printf("%s uses %s on %s for %s %d damage...\n", pok.Name, mv.Name, pok2.Name, mv.Type, mv.Damage)
	pok2.Health -= mv.Damage
	if pok2.Health <= 0 {
		pok2.Health = 0
		fmt.Printf("%s has been defeated!\n", pok2.Name)
		return false, nil
	}
	fmt.Printf("%s has %d hp left!\n", pok2.Name, pok2.Health)
	return true, nil
}
