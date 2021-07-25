package note

import "calendar/internal/app/calendar/domain/date"

type DAO interface {
	Save(n Note) error
	GetByDate(d date.Date) (Note, error)
}
