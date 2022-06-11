package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger *zap.SugaredLogger
}

func New(level string) *Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "../../log_test",
		MaxSize:    100, // megabytes
		MaxBackups: 2,
		MaxAge:     30, // days
	})
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	eLvl := zap.InfoLevel
	err := eLvl.UnmarshalText([]byte(level))
	if err != nil {
		eLvl = zap.InfoLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(w), eLvl),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), eLvl),
	)

	logger := zap.New(core)
	return &Logger{logger.Sugar()}
}

func (l Logger) Info(args ...interface{}) {
	l.Logger.Info(args)
}

func (l Logger) Error(args ...interface{}) {
	l.Logger.Error(args)
}

// TODO
