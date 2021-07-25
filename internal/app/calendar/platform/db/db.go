package db

import (
	"calendar/internal/app/calendar/domain/date"
	"calendar/internal/app/calendar/domain/note"
)

type NoteDAO struct {
}

func NewNoteDAO() *NoteDAO {
	return &NoteDAO{}
}

func (daou *NoteDAO) Save(n note.Note) error {
	return nil
}

func (daou *NoteDAO) GetByDate(d date.Date) (note.Note, error) {
	return note.Note{}, nil
}
