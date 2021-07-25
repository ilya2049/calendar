package zaplogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zapLogger *zap.Logger
}

func (logger Logger) LogInfo(message string, fields ...string) {
	if len(fields) == 0 {
		logger.zapLogger.Info(message)

		return
	}

	logger.zapLogger.Info(message, mapDetailsToZapFields(fields)...)
}

func (logger Logger) LogDebug(message string, fields ...string) {
	if len(fields) == 0 {
		logger.zapLogger.Debug(message)

		return
	}

	logger.zapLogger.Info(message, mapDetailsToZapFields(fields)...)
}

func mapDetailsToZapFields(fields []string) []zapcore.Field {
	fields = balance(fields)

	zapFields := make([]zapcore.Field, 0, len(fields)/2)
	for i := 0; i < len(fields); i += 2 {
		fieldName := fields[i]
		fieldValue := fields[i+1]
		zapFields = append(zapFields, zap.String(fieldName, fieldValue))
	}

	return zapFields
}

func balance(fields []string) []string {
	if len(fields)%2 == 1 {
		fields = append(fields, "")
	}

	return fields
}

func (logger Logger) LogWarn(message string, fields ...string) {
	if len(fields) == 0 {
		logger.zapLogger.Warn(message)

		return
	}

	logger.zapLogger.Warn(message, mapDetailsToZapFields(fields)...)
}

func (logger Logger) LogError(message string, fields ...string) {
	if len(fields) == 0 {
		logger.zapLogger.Error(message)

		return
	}

	logger.zapLogger.Error(message, mapDetailsToZapFields(fields)...)
}
