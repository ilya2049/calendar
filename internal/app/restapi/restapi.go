package restapi

import (
	"calendar/internal/app/restapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func New(httpHandlers *handlers.Handlers) *fiber.App {
	restAPI := fiber.New()

	restAPI.Put("/note", httpHandlers.Note.WriteDownForFuture)

	return restAPI
}
