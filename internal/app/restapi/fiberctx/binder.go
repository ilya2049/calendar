package fiberctx

import (
	"calendar/internal/pkg/event"

	"github.com/gofiber/fiber/v2"
)

const (
	msgInvalidJSONSchema = "invalid json schema"
)

type ErrInvalidJSONSchema struct {
	Cause error
}

func (e *ErrInvalidJSONSchema) Error() string {
	return msgInvalidJSONSchema
}

type Binder struct {
	eventBus event.Bus
}

func NewBinder(eventBus event.Bus) *Binder {
	return &Binder{
		eventBus: eventBus,
	}
}

func (b *Binder) BindJSON(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		eventError := &ErrInvalidJSONSchema{Cause: err}
		b.eventBus.Publish(eventError)
		return eventError
	}

	return nil
}
