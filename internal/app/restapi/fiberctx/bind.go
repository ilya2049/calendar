package fiberctx

import (
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

func BindJSON(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return &ErrInvalidJSONSchema{Cause: err}
	}

	return nil
}
