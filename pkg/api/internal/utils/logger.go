package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitLogger(config ServerConfig) (*zap.Logger, error) {

	logLevel := getLogLevel(config.LogLevel)

	logger, err := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(logLevel),
		OutputPaths: []string{"stdout", config.LogFile},
	}.Build()

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func getLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
		case "info":
			return zapcore.InfoLevel
		case "warn":
			return zapcore.WarnLevel
		case "debug":
			return zapcore.DebugLevel
		case "error":
			return zapcore.ErrorLevel
		case "fatal":
			return zapcore.FatalLevel
		default:
			return zapcore.InfoLevel
	}
}