package fiberctx

import (
	"calendar/internal/pkg/errors"

	stdErrors "errors"

	"github.com/gofiber/fiber/v2"
)

var (
	errInvalidJSON = stdErrors.New("invalid json schema")
)

func BindJSON(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return errors.New(errInvalidJSON).
			SetCode(errors.CodeInvalidInputSchema).
			SetDetail(errors.KeyCause, err.Error())
	}

	return nil
}
