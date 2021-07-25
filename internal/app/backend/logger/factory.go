package zaplogger

import (
	"calendar/internal/pkg/log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// The reason for ignoring is described here: https://github.com/uber-go/zap/issues/772
	ignoredSyncErrorMessage = "sync /dev/stderr: invalid argument"
)

func New(config Config) (logger *Logger, cleanup func(), err error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zapConfig.DisableCaller = true
	zapConfig.DisableStacktrace = true
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel))
	zapLogger, err := zapConfig.Build()

	if err != nil {
		return nil, nil, err
	}

	cleanup = func() {
		if err := zapLogger.Sync(); err != nil && err.Error() != ignoredSyncErrorMessage {
			log.System("can't sync zap logger: " + err.Error())
		}
	}

	return &Logger{zapLogger: zapLogger}, cleanup, nil
}
