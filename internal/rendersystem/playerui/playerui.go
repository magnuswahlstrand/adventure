package playerui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/logger"
	"github.com/kyeett/adventure/internal/rendersystem"
	"github.com/kyeett/adventure/internal/rendersystem/drawutil"
	"github.com/kyeett/adventure/internal/unit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ rendersystem.System = &PlayerUI{}

type PlayerUI struct {
	player *unit.Player
	logger *zap.SugaredLogger
}

func (s *PlayerUI) Draw(screen *ebiten.Image) {
	for i, item := range s.player.Inventory.Inventory {
		drawutil.DrawSprite(screen, comp.PP(i+10,13), item.GetSprite().Image)
	}
}

func NewRender(logLevel zapcore.Level, player *unit.Player) *PlayerUI {
	return &PlayerUI{
		logger:   logger.NewNamed("player", logLevel, logger.BrightBlack),
		player: player,
	}
}

func (s *PlayerUI) Add(_ interface{}) {}

func (s *PlayerUI) Remove(_ comp.ID) {}
