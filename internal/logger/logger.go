package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger     *zap.Logger
	loggerOnce sync.Once
)

// GetLogger returns the singleton zap.Logger instance for the application.
func GetLogger() *zap.Logger {
	loggerOnce.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.TimeKey = "timestamp"

		var err error
		logger, err = cfg.Build()
		if err != nil {
			panic(err)
		}
	})
	return logger
}
