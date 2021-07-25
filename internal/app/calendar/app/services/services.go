package services

import "calendar/internal/app/calendar/domain/note"

type NoteService interface {
	WriteDownForFuture(n note.Note) error
}
