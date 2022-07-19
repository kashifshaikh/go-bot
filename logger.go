package bot

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg *Config) *zap.SugaredLogger {
	var logger *zap.Logger
	var err error
	if cfg.Env == Development {
		logger, err = zap.NewDevelopment()
		// logger, _ = zapDevLogger()
	} else {
		logger, err = zap.NewProduction()
		// logger, _ = zapProdLogger()
	}
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	// Return sugar'd logger. Callers can convert back to plain by Desugar()ing this logger.
	return logger.Sugar()
}

func zapDevLogger() (*zap.Logger, error) {
	return zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()
}

func zapProdLogger() (*zap.Logger, error) {
	return zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()
}
