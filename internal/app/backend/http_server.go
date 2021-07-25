package backend

import (
	"calendar/internal/pkg/log"

	"context"

	"github.com/gofiber/fiber/v2"
)

func runHTTPServer(ctx context.Context, restAPI *fiber.App, listeningAddress string) {
	go func() {
		<-ctx.Done()

		if err := restAPI.Shutdown(); err != nil {
			log.System("can't shutdown the http server: " + err.Error())
		}
	}()

	if err := restAPI.Listen(listeningAddress); err != nil {
		log.System("an http server runtime error: " + err.Error())
	}

	log.System("the http server successfully stopped")
}
