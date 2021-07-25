package note

import (
	"calendar/internal/app/calendar/domain/date"

	stdErrors "errors"
)

var (
	errCantLeaveInPast = stdErrors.New("can't leave a note in the past")
)

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
		return errCantLeaveInPast
	}

	return s.noteDAO.Save(n)
}
