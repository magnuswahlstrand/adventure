package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/game"
	"log"
)

func main() {
	g := game.NewGame()
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
