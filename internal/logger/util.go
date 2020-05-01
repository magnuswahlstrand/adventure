package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
	"time"
)


func NewNamed(name string, level zapcore.Level) *zap.SugaredLogger {
	conf := zap.NewDevelopmentConfig()
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
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