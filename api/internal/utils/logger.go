package utils

import (
	"fmt"

	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	logger, err := zap.NewDevelopmentConfig().Build()
	if err != nil {
		panic(fmt.Errorf("failed to initialize logger: %w", err))
	}
	zap.ReplaceGlobals(logger)
	return logger
}
