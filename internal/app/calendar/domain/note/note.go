package note

import "calendar/internal/app/calendar/domain/date"

type Note struct {
	Date date.Date `json:"date"`
	Text string    `json:"text"`
}
