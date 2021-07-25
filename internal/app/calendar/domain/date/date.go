package date

import (
	"time"
)

const (
	Layout = "2006-01-02"
)

type Date string

func (date Date) ToTime() time.Time {
	t, _ := time.Parse(Layout, string(date))
	return t
}

func IsInFuture(w Watch, date Date) bool {
	return date.ToTime().After(w.Now())
}
