package comp

type Position struct {
	X, Y int
}

func P(x, y int) *Position {
	return &Position{x,y}
}

func (p *Position) GetPosition() *Position {
	return p
}

func (p *Position) Equals(p2 *Position) bool {
	return p.X == p2.X && p.Y == p2.Y
}


