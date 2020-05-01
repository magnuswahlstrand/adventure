package comp

type Hitpoints struct {
	amount int
}

func (p *Hitpoints) GetHitpoints() *Hitpoints {
	return p
}

func HP(amount int) *Hitpoints {
	return &Hitpoints{amount: amount}
}
