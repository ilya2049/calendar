package handlers

import (
	"github.com/gofiber/fiber/v2"

	"calendar/internal/app/calendar/app/services"
	"calendar/internal/app/calendar/domain/note"
	"calendar/internal/app/restapi/fiberctx"
)

type NoteHandler struct {
	noteService services.NoteService
	binder      *fiberctx.Binder
}

func NewNoteHandler(noteService services.NoteService, binder *fiberctx.Binder) *NoteHandler {
	return &NoteHandler{
		noteService: noteService,
		binder:      binder,
	}
}

func (h *NoteHandler) WriteDownForFuture(c *fiber.Ctx) error {
	noteRequest := note.Note{}

	if err := h.binder.BindJSON(c, &noteRequest); err != nil {
		return err
	}

	return h.noteService.WriteDownForFuture(noteRequest)
}
