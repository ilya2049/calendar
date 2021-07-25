package note

import (
	"calendar/internal/app/calendar/domain/date"
)

const (
	msgCantLeaveInPast = "can't leave a note in the past"
)

type ErrCantLeaveInPast struct {
	DateInPast string
}

func (e *ErrCantLeaveInPast) Error() string {
	return msgCantLeaveInPast
}

type Service struct {
	noteDAO DAO
	watch   date.Watch
}

func NewService(noteDAO DAO, watch date.Watch) *Service {
	return &Service{
		noteDAO: noteDAO,
		watch:   watch,
	}
}

func (s *Service) WriteDownForFuture(n Note) error {
	if !date.IsInFuture(s.watch, n.Date) {
		return &ErrCantLeaveInPast{DateInPast: string(n.Date)}
	}

	return s.noteDAO.Save(n)
}
