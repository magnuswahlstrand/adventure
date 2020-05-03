package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/game"
	"log"
)

func main() {
	g := game.NewDefaultGame(1)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowSize(640,480)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
