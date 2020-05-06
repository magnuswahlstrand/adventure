package background

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/logger"
	"github.com/kyeett/adventure/internal/rendersystem"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ rendersystem.System = &Background{}

type Background struct {
	background *ebiten.Image
	logger     *zap.SugaredLogger
}

func (s *Background) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(2,2)
	screen.DrawImage(s.background, opt)
}

func NewRender(logLevel zapcore.Level, background *ebiten.Image) *Background {
	return &Background{
		logger:     logger.NewNamed("background", logLevel, logger.Green),
		background: background,
	}
}

func (s *Background) Add(_ interface{}) {}

func (s *Background) Remove(_ comp.ID) {}
