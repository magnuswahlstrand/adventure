package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
	"time"
)

// Foreground colors.
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	BrightBlack Color = iota + 90
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

// Color represents a text color.
type Color uint8

func NewNamed(name string, level zapcore.Level, color Color) *zap.SugaredLogger {
	conf := zap.NewDevelopmentConfig()
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	conf.EncoderConfig.EncodeName = func(s string, enc zapcore.PrimitiveArrayEncoder) {
		// Colors
		s = fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(color), s)
		zapcore.FullNameEncoder(s, enc)
	}
	conf.EncoderConfig.EncodeTime = func(time.Time, zapcore.PrimitiveArrayEncoder) {}
	conf.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		path := caller.TrimmedPath()
		s := strings.Split(path, "/")
		enc.AppendString(fmt.Sprintf("%24s", s[len(s)-1]))
	}

	conf.Level = zap.NewAtomicLevelAt(level)
	logger, err := conf.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger.Sugar().Named(name)
}

