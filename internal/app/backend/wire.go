package backend

import (
	"calendar/internal/app/backend/logger"
	"calendar/internal/app/calendar/app"
	"calendar/internal/app/calendar/domain/note"
	"calendar/internal/app/calendar/platform/db"
	"calendar/internal/app/calendar/platform/stdwatch"
	"calendar/internal/app/restapi"
	"calendar/internal/app/restapi/fiberctx"
	"calendar/internal/app/restapi/handlers"
	"calendar/internal/pkg/event"
)

func newApplications() (*applications, error) {
	lg, cleanup, err := logger.New(logger.Config{})
	if err != nil {
		return nil, err
	}

	defer cleanup()

	eventBus := event.NewRegistry().
		RegistrerHandlerFunc(lg.LogErrInvalidJSONSchema).
		RegistrerHandlerFunc(lg.LogErrErrCantLeaveNoteInPast).
		Bus()

	fiberBinder := fiberctx.NewBinder(eventBus)

	noteDAO := db.NewNoteDAO()
	noteService := note.NewService(noteDAO, stdwatch.New(), eventBus)
	calendarNoteService := app.NewNoteService(noteService)
	noteHandler := handlers.NewNoteHandler(calendarNoteService, fiberBinder)
	restAPI := restapi.New(handlers.New(noteHandler))

	return &applications{
		restAPI: restAPI,
		config: config{
			restAPIAddress: ":8008",
		},
	}, nil
}
