package backend

import (
	"calendar/internal/app/calendar/app"
	"calendar/internal/app/calendar/domain/note"
	"calendar/internal/app/calendar/platform/db"
	"calendar/internal/app/calendar/platform/stdwatch"
	"calendar/internal/app/restapi"
	"calendar/internal/app/restapi/handlers"
)

func newApplications() (*applications, error) {
	noteDAO := db.NewNoteDAO()
	noteService := note.NewService(noteDAO, stdwatch.New())
	calendarNoteService := app.NewNoteService(noteService)
	noteHandler := handlers.NewNoteHandler(calendarNoteService)
	restAPI := restapi.New(handlers.New(noteHandler))

	return &applications{
		restAPI: restAPI,
		config: config{
			restAPIAddress: ":8008",
		},
	}, nil
}
