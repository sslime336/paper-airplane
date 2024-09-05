package logging

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func Init(logPath, logFile string, devMode bool) {
	if err := os.MkdirAll(logPath, fs.ModeDir); err != nil {
		log.Fatalf("create log dir failed: %v", err)
	}
	logUrl := filepath.Join(logPath, logFile)
	_, err := os.OpenFile(logUrl, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("could not open log file: %v", err)
	}

	encoderConf := zap.NewDevelopmentEncoderConfig()
	generalLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.DebugLevel
	})
	if !devMode {
		encoderConf = zap.NewProductionEncoderConfig()
		generalLevel = zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
			return lev >= zap.InfoLevel
		})
	}

	prepareEncoderConf(&encoderConf)
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConf),
		zapcore.AddSync(os.Stdout),
		generalLevel,
	)
	generalSyncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logUrl,
		MaxSize:    1024, // 1GiB
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
	})
	errorSyncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logPath, "error.log"),
		MaxSize:    500,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   false,
	})
	// Disable colorful logging.
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
	generalCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConf),
		zapcore.AddSync(generalSyncWriter),
		generalLevel,
	)
	errorCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConf),
		zapcore.AddSync(errorSyncWriter),
		zap.NewAtomicLevelAt(zap.ErrorLevel),
	)
	logger = zap.New(zapcore.NewTee(consoleCore, generalCore, errorCore), zap.AddCaller())
}

func Logger() *zap.Logger {
	return logger
}

func Named(name string) *zap.Logger {
	return logger.Named(name)
}

func prepareEncoderConf(encoderConfig *zapcore.EncoderConfig) {
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.NameKey = "name"
	encoderConfig.CallerKey = "caller"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
}
