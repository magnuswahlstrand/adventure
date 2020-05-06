package application

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/adventure/internal/game"
	"log"
	"sync"
)

const (
	StateStarted = "started"
	StateLoading = "loading"

	ModeSingle          = "single"
	ModeHotseat         = "hotseat"
	ModeHostMultiplayer = "host_multiplayer"
	ModeJoinMultiplayer = "join_multiplayer"
)

var _ ebiten.Game = &Application{}

type Application struct {
	game  *game.Game
	state string
	mode  string
	lock  *sync.RWMutex
}

func (a *Application) Update(_ *ebiten.Image) error {
	a.lock.RLock()
	state := a.state
	a.lock.RUnlock()

	if state != StateStarted {
		fmt.Println("loading")
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		a.newGame(ModeSingle)
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		a.newGame(ModeHotseat)
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.Key8) {
		a.newGame(ModeHostMultiplayer)
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.Key9) {
		a.newGame(ModeJoinMultiplayer)
		return nil
	}

	if err := a.game.Update(); err != nil {
		switch err {
		case game.ErrorEndOfGame:
			a.newGame(a.mode)
			return nil
		default:
			return err
		}
	}

	return nil
}

func (a *Application) newGame(mode string) {
	a.lock.RLock()
	a.state = StateLoading
	a.mode = mode
	a.lock.RUnlock()

	go func() {

		switch a.mode {
		case ModeSingle:
			a.game = game.NewDefaultGame(1)
		case ModeHotseat:
			a.game = game.NewDefaultGame(2)
		case ModeHostMultiplayer:
			a.game = game.NewWebsocketGame(true)
		case ModeJoinMultiplayer:
			a.game = game.NewWebsocketGame(false)
		default:
			log.Fatal("invalid mode")
		}

		a.lock.Lock()
		a.state = StateStarted
		a.lock.Unlock()
	}()
}

func (a *Application) Draw(screen *ebiten.Image) {
	a.game.Draw(screen)
}

func (a *Application) drawGameFinished(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "game complete! press R to restart", 10, 2*70)
}

func (a Application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func New() *Application {
	g := game.NewWebsocketGame(true)
	return &Application{g, StateStarted, ModeHostMultiplayer, &sync.RWMutex{}}
}
