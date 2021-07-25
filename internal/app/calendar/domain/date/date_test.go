package date_test

import (
	"calendar/internal/app/calendar/domain/date"
	"calendar/internal/pkg/std"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate_ToTime(t *testing.T) {
	d := date.Date("2014-08-01")

	stdTime := d.ToTime()

	assert.Equal(t, "2014-08-01", stdTime.Format(date.Layout))
}

func TestIsInFuture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		d       date.Date
		nowFunc func() time.Time
		want    bool
	}{
		{
			name: "before",
			d:    date.Date("2014-04-01"),
			nowFunc: func() time.Time {
				return std.NewDate(2014, time.August, 1)
			},
			want: false,
		},
		{
			name: "today",
			d:    date.Date("2014-08-01"),
			nowFunc: func() time.Time {
				return std.NewDate(2014, time.August, 1)
			},
			want: false,
		},
		{
			name: "after",
			d:    date.Date("2014-08-02"),
			nowFunc: func() time.Time {
				return std.NewDate(2014, time.August, 1)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			watch := date.WatchStub{
				NowFunc: tt.nowFunc,
			}

			assert.Equal(t, tt.want, date.IsInFuture(&watch, tt.d))
		})
	}
}
