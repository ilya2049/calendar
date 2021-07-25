package fiberctx

import (
	stdErrors "errors"

	"github.com/gofiber/fiber/v2"
)

var (
	errInvalidJSON = stdErrors.New("invalid json schema")
)

func BindJSON(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return errInvalidJSON
	}

	return nil
}
