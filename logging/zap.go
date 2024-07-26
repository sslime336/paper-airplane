package logging

import (
	"log"
	"os"
	"paper-airplane/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init() {
	localLogPath := config.Log().Path

	serverMode := os.Getenv("AIRP_MODE")
	development := serverMode == "release"

	lv := zapcore.InfoLevel
	if development {
		lv = zapcore.DebugLevel
	}

	logLevel := zap.NewAtomicLevelAt(lv)

	var zc = zap.Config{
		Level:             logLevel,
		Development:       development,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      []string{"stdout", localLogPath},
		ErrorOutputPaths: []string{"stderr", localLogPath},
		InitialFields:    map[string]any{"app": "qqbot-paper-airplane"},
	}

	var err error
	if logger, err = zc.Build(zap.AddCallerSkip(1)); err != nil {
		log.Fatal("init logger failed:", err)
	}
}

func Logger() *zap.Logger {
	return logger
}