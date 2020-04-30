package comp

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
	"log"
)

var (
	SpritePlayer, SpriteEnemySnake *Sprite
)

func init() {
	src := gfx.MustOpenImage("assets/images/tile_0004.png")
	var err error
	playerImage, err := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	SpritePlayer = &Sprite{playerImage}

	src = gfx.MustOpenImage("assets/images/tile_0014.png")
	snakeImage, err := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	SpriteEnemySnake = &Sprite{snakeImage}
}


type Sprite struct {
	*ebiten.Image
}

func (s *Sprite) GetSprite() *Sprite {
	return s
}
