package comp

type Inventory struct {
	Inventory []Inventorable
}

type Inventorable interface{
	GetEntity() Entity
	GetSprite() *Sprite
}

func (i *Inventory) GetInventory() *Inventory {
	return i
}

func (i *Inventory) HasType(t Type) bool {
	for _, item := range i.Inventory {
		if item.GetEntity().Type == t {
			return true
		}
	}

	return false
}

func (i *Inventory) GetType(t Type) Inventorable {
	for _, item := range i.Inventory {
		if item.GetEntity().Type == t {
			return item
		}
	}

	return nil
}

func (i *Inventory) Add(item Inventorable) {
	i.Inventory = append(i.Inventory, item)
}

func (i *Inventory) Remove(item Inventorable) {
	var inv []Inventorable
	for _, inventoryItem := range i.Inventory {
		if inventoryItem.GetEntity().ID == item.GetEntity().ID {
			continue
		}
		inv = append(inv, inventoryItem)
	}
	i.Inventory = inv
}