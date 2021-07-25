package note

import (
	"calendar/internal/app/calendar/domain/date"
	"calendar/internal/pkg/event"
)

const (
	msgCantLeaveInPast = "can't leave a note in the past"
)

type ErrCantLeaveInPast struct {
	DateInPast  string
	CurrentDate string
}

func (e *ErrCantLeaveInPast) Error() string {
	return msgCantLeaveInPast
}

type Service struct {
	eventBus event.Bus
	noteDAO  DAO
	watch    date.Watch
}

func NewService(noteDAO DAO, watch date.Watch, eventBus event.Bus) *Service {
	return &Service{
		noteDAO:  noteDAO,
		watch:    watch,
		eventBus: eventBus,
	}
}

func (s *Service) WriteDownForFuture(n Note) error {
	if !date.IsInFuture(s.watch, n.Date) {
		eventError := &ErrCantLeaveInPast{
			CurrentDate: s.watch.Now().Format(date.Layout),
			DateInPast:  string(n.Date),
		}
		s.eventBus.Publish(eventError)
		return eventError
	}

	return s.noteDAO.Save(n)
}
