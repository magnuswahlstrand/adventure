package controller

import "log"

type Grid struct {
	g map[pos]Describer
}

func (g *Grid) At(x, y int) Describer {
	t, found := g.g[pos{x, y}]
	if !found {
		return theWall
	}
	return t
}

func (g *Grid) Set(x, y int, d Describer) {
	g.g[pos{x, y}] = d
}

func NewGrid(rows ...string) (*Grid, error) {
	g := &Grid{
		map[pos]Describer{},
	}
	for y, row := range rows {
		for x, c := range row {
			var d Describer

			switch c {
			case '-':
				d = theNothing
			case 'x':
				d = theWall
			default:
				log.Fatalf("invalid type: %v", c)
			}

			g.Set(x, y, d)
		}
	}
	return g, nil
}
