package comp

type Hitpoints struct {
	Amount int64
	Max    int64
}

func (p *Hitpoints) GetHitpoints() *Hitpoints {
	return p
}

func HP(amount int64) *Hitpoints {
	return &Hitpoints{Amount: amount, Max: amount}
}
