package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/tiledutil"
	"github.com/lafriks/go-tiled"
)

type G struct {
	*ebiten.Image
}

func (g *G) Update(screen *ebiten.Image) error {
	screen.DrawImage(g.Image, &ebiten.DrawImageOptions{})
	return nil
}

func (g *G) Layout(w, h int) (int, int) {
	return w/4, h/4
}

func main() {
	//b := assets.MustAsset("assets/maps/world.tmx")
	//r := bytes.NewReader(b)
	//m, err := tiled.LoadFromReader("assets/maps/", r)
	//if err != nil {
	//	log.Fatal("load", err)
	//}

	m := tiledutil.MustNewFromFile("assets/maps/world.tmx")


	onlyWalls := func(l *tiled.Layer, tile *tiled.LayerTile, typ string) bool {
		return typ == "wall"
	}
	layer := m.Positions(onlyWalls)
	fmt.Println(layer)
	fmt.Println(len(layer))
	ts := m.Tiles(onlyWalls)
	for _, t := range ts {
		fmt.Println(t.TilesetTile)
		fmt.Println(t.LayerTile)
	}
	//img, err := m.ImageFromLayer("walls")
	//if err != nil {
	//	log.Fatal("new image", err)
	//}
	//
	//
	//backgroundImage := drawutil.MustLoadFromFile("assets/maps/world.tmx", "walls")
	//
	//eImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	//if err != nil {
	//	log.Fatal("new ebiten image", err)
	//}
	//
	//g := &G{
	//	eImg,
	//}
	//ebiten.RunGame(g)
}
