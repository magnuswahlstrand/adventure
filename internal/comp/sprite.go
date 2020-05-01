package comp

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/assets"
	"github.com/peterhellberg/gfx"
	"log"
)

var (
	SpritePlayer, SpriteEnemySnake, SpriteEnemyRat, SpriteChest, SpriteKey *Sprite
)

func init() {
	SpritePlayer = &Sprite{loadImageOrFatal("assets/images/tile_0004.png")}
	SpriteEnemySnake = &Sprite{loadImageOrFatal("assets/images/tile_0014.png")}
	SpriteEnemyRat = &Sprite{loadImageOrFatal("assets/images/tile_0016.png")}
	SpriteChest = &Sprite{loadImageOrFatal("assets/images/tile_0039.png")}
	SpriteKey = &Sprite{loadImageOrFatal("assets/images/tile_0057.png")}
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
