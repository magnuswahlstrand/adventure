package comp

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure/assets"
	"github.com/peterhellberg/gfx"
	"log"
)

var (
	SpriteWall, SpritePlayer, SpritePlayer2, SpriteEnemySnake, SpriteEnemyRat, SpriteChest, SpriteKey, SpriteGoal, SpriteDoor *Sprite
)

func init() {
	SpriteWall = &Sprite{loadImageOrFatal("assets/images/tile_0001.png")}
	SpritePlayer = &Sprite{loadImageOrFatal("assets/images/tile_0004.png")}
	SpritePlayer2 = &Sprite{loadImageOrFatal("assets/images/tile_0006.png")}
	SpriteEnemySnake = &Sprite{loadImageOrFatal("assets/images/tile_0014.png")}
	SpriteEnemyRat = &Sprite{loadImageOrFatal("assets/images/tile_0016.png")}
	SpriteDoor = &Sprite{loadImageOrFatal("assets/images/tile_0024.png")}
	SpriteChest = &Sprite{loadImageOrFatal("assets/images/tile_0039.png")}
	SpriteKey = &Sprite{loadImageOrFatal("assets/images/tile_0057.png")}
	SpriteGoal = &Sprite{loadImageOrFatal("assets/images/tile_0081.png")}
}

func loadImageOrFatal(path string) *ebiten.Image {
	b, err := assets.Asset(path)
	if err != nil {
		log.Fatal(err)
	}

	src, err := gfx.DecodeImageBytes(b)
	if err != nil {
		log.Fatal(err)
	}

	img, err := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

type Sprite struct {
	*ebiten.Image
}

func (s *Sprite) GetSprite() *Sprite {
	return s
}
