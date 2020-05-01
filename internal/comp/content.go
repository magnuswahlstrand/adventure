package comp

type Content struct {
	Type string
}

func (c *Content) GetContent() *Content {
	return c
}