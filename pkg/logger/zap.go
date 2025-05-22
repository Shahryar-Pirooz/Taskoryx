package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Init(isProduction bool) {
	once.Do(func() {
		var cfg zap.Config

		if isProduction {
			cfg = zap.NewProductionConfig()
		} else {
			cfg = zap.NewDevelopmentConfig()
		}

		var err error

		logger, err = cfg.Build()
		if err != nil {
			panic("Failed to initialize logger: " + err.Error())
		}
		zap.ReplaceGlobals(logger)
	})
}

func Get() *zap.Logger {
	Init(false)
	return logger
}
