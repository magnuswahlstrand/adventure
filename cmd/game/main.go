package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure/internal/application"
	"log"
)

func main() {
	a := application.New()
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowSize(640,480)
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}
