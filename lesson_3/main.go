package main

import (
	//pok "github.com/helenraven/pockemon"
	pok "github.com/helenraven/pockemon/v2"
)

func main() {
	p1 := &pok.Pokemon{"Bulbasaur"}
	p2 := &pok.Pokemon{"Pikachu"}

	p1.IChooseU()
	p2.IChooseU()

	pok.StartFight(p1, p2)
	pok.Winner(p1, p2)
}
