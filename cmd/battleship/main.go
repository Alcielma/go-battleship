package main

import (
	"log"

	"gobattleship/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Battleship - Test")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
