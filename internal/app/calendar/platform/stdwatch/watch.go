package stdwatch

import "time"

func New() *Watch {
	return &Watch{}
}

type Watch struct{}

func (*Watch) Now() time.Time {
	return time.Now()
}
