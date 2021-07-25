package backend

import "github.com/gofiber/fiber/v2"

type applications struct {
	restAPI *fiber.App

	config config
}

type config struct {
	restAPIAddress string
}
