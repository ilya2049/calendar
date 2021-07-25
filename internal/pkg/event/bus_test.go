package event_test

import (
	"calendar/internal/pkg/event"

	"testing"

	"github.com/stretchr/testify/assert"
)

type counter struct {
	eventCount int
}

type countEvent struct{}

func (c *counter) count(ev countEvent) {
	c.eventCount++
}

func Test_bus_Publish(t *testing.T) {
	c := counter{}

	eventBus := event.NewRegistry().
		RegistrerHandlerFunc(c.count).
		Bus()

	eventBus.Publish(countEvent{})
	eventBus.Publish(countEvent{})
	eventBus.Publish(countEvent{})

	assert.Equal(t, 3, c.eventCount)
}
