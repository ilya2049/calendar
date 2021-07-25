package logger

import (
	"calendar/internal/app/calendar/domain/note"
	"calendar/internal/app/restapi/fiberctx"

	"go.uber.org/zap"
)

type Logger struct {
	zapLogger *zap.Logger
}

func (lg *Logger) LogErrInvalidJSONSchema(evnt *fiberctx.ErrInvalidJSONSchema) {
	lg.zapLogger.Info(evnt.Error(),
		zap.String(keyCause, evnt.Cause.Error()),
	)
}

func (lg *Logger) LogErrErrCantLeaveNoteInPast(evnt *note.ErrCantLeaveInPast) {
	lg.zapLogger.Info(evnt.Error(),
		zap.String(keyDateInPast, evnt.DateInPast),
		zap.String(keyCurrentDate, evnt.CurrentDate),
	)
}
