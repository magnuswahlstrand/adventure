package system

// A RenderSystem is updated every iteration, and draws to a screen
type System interface {
	Add(v interface{})
}
