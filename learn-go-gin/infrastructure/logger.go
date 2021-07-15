package infrastructure

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

// NewLogger sets up logger
func NewLogger(env Env) Logger {
	config := zap.NewDevelopmentConfig()

	if env.ENVIRONMENT == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if env.ENVIRONMENT == "production" && env.LOG_OUTPUT != "" {
		// config.OutputPaths = []string{env.LogOutput}
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := config.Build()

	sugar := logger.Sugar()

	return Logger{
		SugaredLogger: sugar,
	}
}
