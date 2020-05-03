package translation

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/event"
	"github.com/kyeett/single-player-game/internal/inputhandler"
	"github.com/kyeett/single-player-game/internal/logger"
	"github.com/kyeett/single-player-game/internal/system"
	"github.com/kyeett/single-player-game/internal/unit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ system.System = &Translation{}
var _ inputhandler.InputHandler = &Translation{}

type Translation struct {
	entities   map[comp.ID]TranslatableEntity
	logger     *zap.SugaredLogger

	player *unit.Player
}

func NewTranslation(logLevel zapcore.Level, player *unit.Player) *Translation {
	return &Translation{
		entities:   map[comp.ID]TranslatableEntity{},
		logger:     logger.NewNamed("transl", logLevel, logger.Yellow),
		player:     player,
	}
}

type Translatable interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
}

type TranslatableEntity struct {
	comp.Entity
	*comp.Position
	source interface{}
}

func (s *TranslatableEntity) GetSource() interface{} {
	return s.source
}

func (s *Translation) Update(_ event.Event) []*command.Command {
	return nil
}

func (s *Translation) Add(v interface{}) {
	i, ok := v.(Translatable)
	if !ok {
		return
	}
	e := TranslatableEntity{
		Entity:   i.GetEntity(),
		Position: i.GetPosition(),
		source:   v,
	}
	s.logger.Info("entity added")
	s.entities[e.ID] = e
}

func (s *Translation) Remove(id comp.ID) {
	delete(s.entities, id)
}

func (s *Translation) MoveBy(dx, dy int) *comp.Position {
	pos := s.player.Position
	x0, y0 := pos.X, pos.Y
	return comp.PP(x0+dx, y0+dy)
}

func (s *Translation) GetEvent() event.Event {
	// Interpret events
	var mv *comp.Position
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		mv = s.MoveBy(1, 0)
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		mv = s.MoveBy(-1, 0)
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		mv = s.MoveBy(0, -1)
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		mv = s.MoveBy(0, 1)
	default:
		return nil
	}

	//actor, err := s.findByID(v.ActorID)
	target := s.findAtPosition(mv)
	switch target.Type {
	case comp.TypeEnemy,
		comp.TypeChest,
		comp.TypeItem,
		comp.TypeDoor,
		comp.TypeGoal:
		return s.playerInteract(target)
	case comp.TypeNil:
		return event.Move{
			Actor:    s.player.ID,
			Position: *mv,
		}
	default:
		return nil
	}
}
