package date

import "time"

type Watch interface {
	Now() time.Time
}

type WatchStub struct {
	NowFunc func() time.Time
}

func (w *WatchStub) Now() time.Time {
	return w.NowFunc()
}
