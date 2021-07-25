package app

import (
	"calendar/internal/app/calendar/domain/note"
)

type NoteService struct {
	noteService *note.Service
}

func NewNoteService(noteService *note.Service) *NoteService {
	return &NoteService{
		noteService: noteService,
	}
}

func (s *NoteService) WriteDownForFuture(n note.Note) error {
	return s.noteService.WriteDownForFuture(n)
}
