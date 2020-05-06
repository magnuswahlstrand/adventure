module github.com/kyeett/adventure

go 1.14

require (
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hajimehoshi/ebiten v1.11.0
	github.com/kyeett/tiledutil v0.0.0-20200506190859-9964784d4c41
	github.com/lafriks/go-tiled v0.0.0-20200408083348-2d002f58d28d
	github.com/peterhellberg/gfx v0.0.0-20200413155834-f94ca2597d91
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.15.0
	nhooyr.io/websocket v1.8.5
)

// replace github.com/kyeett/tiledutil => ../tiledutil
