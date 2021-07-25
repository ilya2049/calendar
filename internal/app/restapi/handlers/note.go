package handlers

import (
	"github.com/gofiber/fiber/v2"

	"calendar/internal/app/calendar/app/services"
	"calendar/internal/app/calendar/domain/note"
	"calendar/internal/app/restapi/fiberctx"
)

type NoteHandler struct {
	noteService services.NoteService
}

func NewNoteHandler(noteService services.NoteService) *NoteHandler {
	return &NoteHandler{
		noteService: noteService,
	}
}

func (h *NoteHandler) WriteDownForFuture(c *fiber.Ctx) error {
	noteRequest := note.Note{}

	if err := fiberctx.BindJSON(c, &noteRequest); err != nil {
		return err
	}

	return h.noteService.WriteDownForFuture(noteRequest)
}
