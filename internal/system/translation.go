package system

import (
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"go.uber.org/zap"
	"log"
)

var _ System = &Translation{}

type Translation struct {
	entities []SomeInterface
	logger   *zap.Logger
}

func NewTranslation(logger *zap.Logger) *Translation {
	return &Translation{
		entities: []SomeInterface{},
		logger:   logger,
	}
}

type SomeInterface interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
}

const (
	translationSystem = "translationsystem"
)

func (t *Translation) Add(v interface{}) {
	e, ok := v.(SomeInterface)
	if !ok {
		return
	}
	t.logger.Info("add entity to " + translationSystem)
	t.entities = append(t.entities, e)
}

func (t *Translation) target(p *comp.Position) comp.Entity {
	if p.X >= 6 || p.X < 0 {
		return comp.BaseWall
	}

	if p.Y >= 6 || p.Y < 0 {
		return comp.BaseWall
	}

	for _, e := range t.entities {
		if p.Equals(e.GetPosition()) {
			return e.GetEntity()
		}
	}

	return comp.BaseNil
}

func (t *Translation) Translate(c interface{}) []*command.Command {
	if c == nil {
		return nil
	}

	switch v := c.(type) {
	case command.MoveTo:

		target := t.target(v.Target)
		switch target.GetType() {
		case comp.TypeEnemy:
			commands := t.interact(v.Actor.GetEntity(), target)
			return commands
		case comp.TypeNil:
			return []*command.Command{command.MoveToCommand(v.Actor, v.Target)}
		default:
			return []*command.Command{}
		}
	default:
		log.Fatal("translate:", v)
	}
	return nil
}

func (t *Translation) interact(e, target comp.Entity) []*command.Command {
	switch e.Type {
	case comp.TypePlayer:
		return playerInteract(e, target)
	}

	return nil
}

func playerInteract(e, target comp.Entity) []*command.Command {
	return nil
}
